package Iterator.A1;

import java.util.Iterator;
import java.util.NoSuchElementException;

// Iterator: 特定のコレクションの要素を反復することができるような方法を提供するためのインターフェース
public class BookShelfIterator implements Iterator<Book> {
    private final BookShelf bookShelf;
    private int index;

    // コンストラクタでBookShelfを受け取り、初期化
    public BookShelfIterator(BookShelf bookShelf) {
        this.bookShelf = bookShelf;
        this.index = 0;
    }

    @Override
    public boolean hasNext() {
        return index < bookShelf.getLength();
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
