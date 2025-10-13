import java.util.List;
import java.util.ArrayList;

class Task {
    private final int id;

    public Task(int id) {
        this.id = id;
    }

    public void run() {
        System.out.println("Running task " + id);
    }
}


class Runner {
    public void run(Task task) {
        task.run();
    }

    public void runAll(List<Task> tasks) {
        for (Task task : tasks) {
            run(task);
        }
    }
}

class RunCounter extends Runner {
    private int count = 0;

    @Override
    public void run(Task task) {
        count++;
        super.run(task);
    }

    @Override
    public void runAll(List<Task> tasks) {
        count += tasks.size();
        super.runAll(tasks);
    }

    public int getCount() {
        return count;
    }
}

public class Main {
    public static void main(String[] args) {
        Task t1 = new Task(1);
        Task t2 = new Task(2);
        Task t3 = new Task(3);
        List<Task> tasks = List.of(t1, t2, t3);

        Runner runner = new Runner();
        runner.runAll(tasks);
/*
        RunCounter runner = new RunCounter();
        runner.runAll(tasks);
        System.out.println("Tasks executed: " + runner.getCount());
*/
    }
}