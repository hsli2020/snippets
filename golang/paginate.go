package html

import "fmt"
import "io"
import "net/url"

type Pagination struct {
	URL    url.URL
	Offset int
	Limit  int
	N      int
}

func (r *Pagination) Render(w io.Writer) {
	// Do not render if no limit exists or there is only one page.
	if r.Limit == 0 || r.N <= r.Limit {
		return
	}

	// Determine the page range & current page.
	current := (r.Offset / r.Limit) + 1
	pageN := ((r.N - 1) / r.Limit) + 1

	prev := current - 1
	if prev <= 0 {
		prev = 1
	}
	next := current + 1
	if next >= pageN {
		next = pageN
	}

	// Print container & "previous" link.
	fmt.Fprint(w, `<nav aria-label="Page navigation">`)
	fmt.Fprint(w, `<ul class="pagination pagination-sm justify-content-end mb-0">`)
	fmt.Fprintf(w, `<li class="page-item"><a class="page-link" href="%s">Previous</a></li>`,
		r.pageURL(current-1))

	// Print a button for each page number.
	for page := 1; page <= pageN; page++ {
		className := ""
		if page == current {
			className = " active"
		}
		fmt.Fprintf(w, `<li class="page-item %s"><a class="page-link" href="%s">%d</a></li>`,
			className, r.pageURL(page), page)
	}

	// Print "next" link & close container.
	fmt.Fprintf(w, `<li class="page-item"><a class="page-link" href="%s">Next</a></li>`,
		r.pageURL(current+1))
	fmt.Fprint(w, `</ul></nav>`)
}

func (r *Pagination) pageURL(page int) string {
	// Ensure page number is within min/max.
	pageN := ((r.N - 1) / r.Limit) + 1
	if page < 1 {
		page = 1
	} else if page > pageN {
		page = pageN
	}

	q := r.URL.Query()
	q.Set("offset", fmt.Sprint((page-1)*r.Limit))
	u := url.URL{Path: r.URL.Path, RawQuery: q.Encode()}
	return u.String()
}
