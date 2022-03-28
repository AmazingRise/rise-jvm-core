public class Add {

    public static void main() {
        int result = Add(1, 2);
        int result2 = Calc();
        System.out.println("Hello");
        return;
    }

    //private static String hello = "Hello, ";

    public static int Calc() {
        int a = 100;
        int b = 200;
        int c = 300;
        return (a + b) * c;
    }

    public static int Add(int a, int b) {
        return a + b;
    }
}