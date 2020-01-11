// Greetable.java
public interface Greetable {
    public String greet();
}

// Dog.java
public class Dog implements Greetable {
    public String greet() {
        return "Woof!";
    }

}

    // Runner.java
public static void main(String[] args) {
    Greetable greetable = new Dog();
    greetable.greet();
}