import java.util.Scanner;

public class TestFac {
    public static void main(String[] args) throws Exception {
        SimpleFactory fac = new SimpleFactory();
        Scanner sc = new Scanner(System.in);
        System.out.print("input:");
        String selects = sc.nextLine();
        fac.produce(selects).show();
    }
}
