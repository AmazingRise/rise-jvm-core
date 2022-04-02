public class Overload {

    public static Boolean compare(int a, int b) {
        return a == b;
    }

    public static Boolean compare(Overload a, Overload b) {
        return a == b;
    }

    public static void main(String[] args) {
        Overload a = new Overload();
        Overload b = new Overload();

        System.out.println(compare(a, b));
        System.out.println(compare(a, a));
        System.out.println(compare(1, 2));
        System.out.println(compare(1, 1));
    }
}