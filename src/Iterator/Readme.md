

あるデータの集まり（リスト、配列、ツリーなど）の要素を順番に取り出して処理する際に、データの種類や内部構造に関わらず、統一的な方法でアクセスできるようにするためのデザインパターン

### メリット

- データの種類や内部構造が変わっても、アクセス方法を変更する必要がない
    - データの実装が変わっても、プログラムの他の部分に影響を与えずに対応できる
- 各 Iterator オブジェクトが独自の状態を持つため、同じデータの集まりに対して複数の処理を同時に行うことができる

## 構成要素

### Iterator（イテレータ）

- 要素を順番にスキャンしていくためのインターフェース

### Concrete Iterator（具体的なイテレータ）

- Iterator<Book> インターフェースを実装するクラス
- 集約体の要素を順番にスキャンするためのメソッドを実装する

```java
package Iterator;

import java.util.Iterator;
import java.util.NoSuchElementException;

public class BookShelfIterator implements Iterator<Book> {
    private BookShelf bookShelf;
    private int index;

    public BookShelfIterator(BookShelf bookShelf) {
        this.bookShelf = bookShelf;
        this.index = 0;
    }

    @Override
    public boolean hasNext() {
        if (index < bookShelf.getLength()) {
            return true;
        } else {
            return false;
        }
    }

    @Override
    public Book next() {
        if (!hasNext()) {
            throw new NoSuchElementException();
        }
        Book book = bookShelf.getBookAt(index);
        index++;
        return book;
    }
}
```

### Aggregate（集約体）

```java
package Iterator;

public class Book {
    private String name;

    public Book(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }
}
```

### Concrete Aggregate（具体的な集約体）

- Book オブジェクトの集まりを表すクラス

```java
package Iterator;

import java.util.Iterator;

public class BookShelf implements Iterable<Book> {
    private Book[] books;
    private int last = 0;

    public BookShelf(int maxsize) {
        this.books = new Book[maxsize];
    }

    public Book getBookAt(int index) {
        return books[index];
    }

    public void appendBook(Book book) {
        this.books[last] = book;
        last++;
    }

    public int getLength() {
        return last;
    }

    @Override
    public Iterator<Book> iterator() {
        return new BookShelfIterator(this);
    }
}
```