public class Obj {

    public int number;

    public int GetNum() {
        return this.number;
    }

    public void SetNum(int num) {
        this.number = num;
    }

    public static void main(String[] args) {
        Obj obj = new Obj();
        obj.SetNum(1);
        int num = obj.GetNum();
        System.out.println(num);
    }

}