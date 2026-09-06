// Command quizserver serves the lecture quiz engine: a Go HTTP server with an
// embedded TypeScript web UI, reading quizzes from YAML files.
//
//	go run ./src/quiz/cmd/quizserver -dir src/quiz/quizzes
package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/s-macke/concepts-of-programming-languages/src/quiz/quiz"
	"github.com/s-macke/concepts-of-programming-languages/src/quiz/server"
	"github.com/s-macke/concepts-of-programming-languages/src/quiz/session"
	"github.com/s-macke/concepts-of-programming-languages/src/quiz/web"
)

func main() {
	addr := flag.String("addr", ":8080", "address to listen on")
	dir := flag.String("dir", "src/quiz/quizzes", "directory with quiz YAML files")
	ttl := flag.Duration("ttl", time.Hour, "how long an idle session is kept")
	flag.Parse()

	quizzes, problems := quiz.LoadDir(*dir)
	for _, p := range problems {
		log.Printf("skipping quiz: %v", p)
	}
	log.Printf("loaded %d quizzes from %s", len(quizzes), *dir)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	store := session.NewStore(*ttl)
	go store.Janitor(ctx, time.Minute)

	srv := &http.Server{
		Addr:              *addr,
		Handler:           server.New(quizzes, store, web.Assets()).Routes(),
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("quiz server listening on http://localhost%s", *addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown: %v", err)
	}
}
