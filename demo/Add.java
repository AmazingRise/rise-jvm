public class Add {

    public static void main(String[] args) {
        System.out.println("Hello world!");
        int a = 100;
        int b = 200;
        int c = 300;
        int d = Calc(a, b, c);
        System.out.print("(100+200)*300=");
        System.out.println(d);
        return;
    }

    public static int Calc(int a, int b, int c) {
        return (a + b) * c;
    }
}