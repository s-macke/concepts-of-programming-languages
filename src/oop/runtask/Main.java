import java.util.List;
import java.util.ArrayList;

interface Task {
    void run();
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
        RunCounter runner = new RunCounter();

        Task t1 = () -> System.out.println("Task 1 running");
        Task t2 = () -> System.out.println("Task 2 running");
        Task t3 = () -> System.out.println("Task 3 running");

        List<Task> tasks = new ArrayList<>(t1, t2, t3);
        runner.runAll(tasks);
        System.out.println("Tasks executed: " + runner.getCount());
    }
}