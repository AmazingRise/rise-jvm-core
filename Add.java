public class Add {

    public static void main() {
        int a = 100;
        int b = 200;
        int c = 300;
        int d = Calc(a, b, c);
        return;
    }

    //private static String hello = "Hello, ";

    public static int Calc(int a, int b, int c) {
        return (a + b) * c;
    }

    public static int Add(int a, int b) {
        return a + b;
    }
}