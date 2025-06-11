# ecommerce-go

## 🧠 Understanding Pointers and Constructors in Go (Using `repository_sql.go`)

In Go, understanding how to use `*` (dereference) and `&` (address-of) is essential when working with dependency injection and clean architecture.

---

### 📌 Key Concepts

| Symbol | Meaning                         | Usage                                |
|--------|----------------------------------|--------------------------------------|
| `&`    | Address of a value               | Used when passing a pointer          |
| `*`    | Dereferencing / Pointer receiver | Used in function signatures to work with the real value |

---

### ✅ Constructor Pattern in Go

Go doesn’t support traditional constructors like C# or Java, but by convention we use `NewType()` functions that return initialized struct instances (often as interfaces).

Here’s an example using a **SQL repository** for the `category` entity:

```go
func NewSQLRepository(db *sql.DB) Storage {
    return &repositorySQL{db: db}
}


/*
╔══════════════════════════════════════════════════════════════════╗
║     💡 Go Pointers and Constructor Usage Explanation            ║
╚══════════════════════════════════════════════════════════════════╝

📌 This file defines a SQL-based repository for categories.
   It uses a constructor function: NewSQLRepository(db *sql.DB)

🔹 Pointers:
   - Go functions can receive either values or pointers.
   - If a function or method expects a pointer (*Type), you usually pass it using &.

🔸 Constructor: NewSQLRepository
   - Signature: func NewSQLRepository(db *sql.DB) Storage
   - Parameter: expects a pointer to a sql.DB object.
   - Return: a pointer to a repositorySQL struct that implements the Storage interface.

✅ DO:
   db, _ := sql.Open("mysql", "...") // db is already *sql.DB
   repo := NewSQLRepository(db)      // Correct: db is already a pointer

🚫 DON'T:
   db := sql.DB{}                    // db is a value
   repo := NewSQLRepository(db)      // Error: expected *sql.DB

✅ FIX:
   repo := NewSQLRepository(&db)     // Correct: pass pointer explicitly

📎 Summary:
   - Use `&` when you have a value and need a pointer.
   - Use `*` to define functions that expect or operate on pointers.
   - `sql.Open(...)` returns a *sql.DB, so no `&` is needed in normal usage.

*/
```

---

### 💻 Example Code

```go
package category

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/ddessilvestri/ecommerce-go/models"
)

// This struct acts like a "class" in Go.
// It implements the Storage interface for SQL-based storage.
type repositorySQL struct {
	db *sql.DB // Dependency to the database connection
}

/*
	🔹 C# Equivalent:

	public class CategoryRepository {
	    private SqlConnection _db;
	    public CategoryRepository(SqlConnection db) {
	        _db = db;
	    }
	}
*/

// Constructor-like function (Go does not support constructors like C# or Java).
// By convention, we use New<Name>() to instantiate and return the interface type.
func NewSQLRepository(db *sql.DB) Storage {
	// We return a pointer to the struct instance
	return &repositorySQL{db: db}
}

/*
	🔹 C# Equivalent:

	public interface ICategoryStorage {
	    long InsertCategory(Category c);
	}

	public class CategoryRepository : ICategoryStorage {
	    public long InsertCategory(Category c) {
	        // SQL logic here
	    }
	}
*/

// Method bound to the repositorySQL struct.
// The receiver is a pointer (*repositorySQL), which allows modifying internal state
// and avoids copying the struct on each method call.
func (r *repositorySQL) InsertCategory(c models.Category) (int64, error) {
	// Build a safe SQL INSERT query using the squirrel package
	query, args, err := squirrel.
		Insert("category").
		Columns("Categ_Name", "Categ_Path").
		Values(c.CategName, c.CategPath).
		ToSql()

	if err != nil {
		return 0, err
	}

	// Execute the query with the generated SQL and arguments
	result, err := r.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	// Return the last inserted ID
	return result.LastInsertId()
}
```

---

### 🧪 Pointer Theory in Practice

```go
type User struct {
	Name string
}

func updateName(u *User) {
	u.Name = "Alice"
}

func main() {
	user := User{Name: "Original"}
	updateName(&user) // Pass address to modify original value
}
```

## 🧱 Design Patterns in Use

This project follows a **Clean Architecture** approach and implements several key design patterns to ensure maintainability, testability, and scalability.

### ✅ Patterns Currently Applied

| Pattern                 | Purpose                                                                 | Implementation Example                                              |
|-------------------------|-------------------------------------------------------------------------|----------------------------------------------------------------------|
| **Repository Pattern**   | Abstracts data access to support multiple storage backends              | [`category/repository_sql.go`](./internal/category/repository_sql.go) implements `Storage` interface |
| **Factory Pattern**      | Standardized way to construct components (similar to constructors)      | `NewSQLRepository`, `NewCategoryService`, `NewCategoryHandler`      |
| **Dependency Injection** | Injects dependencies from higher layers to lower layers, improves testing | Passed from `main.go → router → handler → service`                  |
| **Pointer Injection**    | Go idiomatic way to share resources like `*sql.DB` without copying      | Used across `repository`, `service`, `handler` layers               |

---

### ✨ Example: Repository Pattern

```go
type Storage interface {
    InsertCategory(c models.Category) (int64, error)
}

type repositorySQL struct {
    db *sql.DB
}

func NewSQLRepository(db *sql.DB) Storage {
    return &repositorySQL{db: db}
}
```

This design allows you to swap `repositorySQL` with a `MongoRepository` or `MockRepository` without changing the business logic that depends on the interface.

---

### 🔜 Coming Next: Middleware Pattern

We plan to introduce the **Middleware Pattern** to centralize and simplify cross-cutting concerns such as:

- ✅ JWT Token validation
- ✅ Admin access checks (`UserIsAdmin`)
- 🧪 Future logging, panic recovery, etc.

This will ensure cleaner and reusable handler logic.

#### 💡 Planned Branch

```
feature/add-middleware-auth
```

The goal is to enable middleware chaining that wraps handlers like:

```go
handlerWithMiddleware := middleware.Authenticate(UserIsAdmin)(handler.Post)
```