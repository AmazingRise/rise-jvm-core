public class MemoryLeak {
    public int num;

    public static void main(String[] args) {
        int i;
        for (i=0;i<100000;i++) {
            MemoryLeak leak = new MemoryLeak();
            leak.num = i;
        }
    }
}