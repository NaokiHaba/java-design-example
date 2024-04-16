package Iterator.A1;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

// Iterableインターフェースを実装することで、拡張for文を使えるようになる
public class BookShelf implements Iterable<Book> {
    private final List<Book> books;

    // 任意の要素数の本棚を作る
    public BookShelf(int initialsize) {
        this.books = new ArrayList<>(initialsize);
    }

    // 本棚の指定した位置にある本を取得
    public Book getBookAt(int index) {
        return books.get(index);
    }

    // 本棚に本を追加
    public void appendBook(Book book) {
        books.add(book);
    }

    // 本棚の長さを取得
    public int getLength() {
        return books.size();
    }

    // iterator()メソッドがオーバーライドされていることを明示的に示す
    @Override
    public Iterator<Book> iterator() {
        return new BookShelfIterator(this);
    }
}
