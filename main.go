package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	// максимально допустимое число ошибок при парсинге
	errorsLimit = 1000

	// число результатов, которые хотим получить
	resultsLimit = 10000

	// общий таймаут
	maxTime = 2 * time.Second
)

var (
	// адрес в интернете (например, https://en.wikipedia.org/wiki/Lionel_Messi)
	url string

	// насколько глубоко нам надо смотреть (например, 10)
	depthLimit int
)

// Как вы помните, функция инициализации стартует первой
func init() {
	// задаём и парсим флаги
	flag.StringVar(&url, "url", "", "url address")
	flag.IntVar(&depthLimit, "depth", 3, "max depth for run")
	flag.Parse()

	// Проверяем обязательное условие
	if url == "" {
		log.Print("no url set by flag")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	fmt.Printf("Process PID : %v\n", os.Getpid())

	started := time.Now()

	ctx, cancel := context.WithCancel(context.Background())
	go watchSignals(cancel)
	defer cancel()

	crawler := newCrawler(depthLimit)

	watchUsrSignal(ctx, crawler)

	// создаём канал для результатов
	results := make(chan crawlResult)

	// запускаем горутину для чтения из каналов
	done := watchCrawler(ctx, results, errorsLimit, resultsLimit)

	// запуск основной логики
	// внутри есть рекурсивные запуски анализа в других горутинах
	// добавила wait group для корректного оджидания завершения всех горутин
	crawler.wg.Add(1)

	// общий таймаут для парсера
	ctxParse, cancelParse := context.WithTimeout(ctx, maxTime)
	defer cancelParse()
	crawler.run(ctxParse, url, results, 0)

	crawler.wg.Wait()
	close(results)

	// ждём завершения работы чтения в своей горутине
	<-done

	log.Println(time.Since(started))
}

func watchUsrSignal(ctx context.Context, c *crawler) {
	usrSignalChan := make(chan os.Signal)

	signal.Notify(usrSignalChan,
		syscall.SIGUSR1)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case sig := <-usrSignalChan:
				c.Mutex.Lock()
				c.maxDepth += 2
				c.Mutex.Unlock()
				log.Printf("got signal %q, max depth increades by 2", sig.String())
			}
		}
	}()

}

// ловим сигналы выключения
func watchSignals(cancel context.CancelFunc) {
	osSignalChan := make(chan os.Signal)

	signal.Notify(osSignalChan,
		syscall.SIGINT,
		syscall.SIGTERM)

	sig := <-osSignalChan
	log.Printf("got signal %q", sig.String())

	// если сигнал получен, отменяем контекст работы
	cancel()
}

func watchCrawler(ctx context.Context, results <-chan crawlResult, maxErrors, maxResults int) chan struct{} {
	readersDone := make(chan struct{})

	go func() {
		defer close(readersDone)
		for {
			select {
			case <-ctx.Done():
				return

			case result, ok := <-results:
				if !ok {
					return
				}

				if result.err != nil {
					maxErrors--
					if maxErrors <= 0 {
						log.Println("max errors exceeded")
						return
					}
					continue
				}

				if result.msg != "" {
					log.Printf("crawling result: %v", result.msg)
					maxResults--
					if maxResults <= 0 {
						log.Println("got max results")
						return
					}
				}
			}
		}
	}()

	return readersDone
}
