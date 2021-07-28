public class SimpleFactory {
    public FactoryTest produce(String selects) {
        if (selects.equals("Factory1")) {
            return new Factory1();
        } else if (selects.equals("Factory2")) {
            return new Factory2();
        } else {
            System.out.println("No product!");
            return null;
        }
    }
}
