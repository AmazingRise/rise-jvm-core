public class MemoryLeak {
    public int num;

    public static void main(String[] args) {
        int i;
        for (i=0;i<10;i++) {
            MemoryLeak leak = new MemoryLeak();
            leak.num = i;
        }
    }
}