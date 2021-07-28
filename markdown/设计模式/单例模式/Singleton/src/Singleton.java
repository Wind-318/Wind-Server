public class Singleton {
    private static Singleton UniqueSingleton = new Singleton();

    private Singleton() {}
    
    public static Singleton GetSingleton() {
        return UniqueSingleton;
    }

    public void output() {
        System.out.println("Hello Singleton!");
    }
}
