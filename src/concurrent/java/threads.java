import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

public class Main {
    public static void main(String[] args) {
        int numberOfTasks = 1000;
        int threadPoolSize = 1000;

        ExecutorService executorService = Executors.newFixedThreadPool(threadPoolSize);

        for (int i = 0; i < numberOfTasks; i++) {
            executorService.submit(() -> {
                try {
                    Thread.sleep(10000); // Wait for 3 seconds
                    System.out.println("Thread " + Thread.currentThread().getId() + " finished waiting.");
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            });
        }

        executorService.shutdown();

        try {
            // Optional: Wait for all tasks to finish execution
            executorService.awaitTermination(Long.MAX_VALUE, TimeUnit.NANOSECONDS);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}