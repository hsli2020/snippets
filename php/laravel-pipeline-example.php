<?php # intro Pipeline

use Illuminate\Support\Facades\Pipeline;

$comment = "  Hey you :-) ";

$result = Pipeline::send($comment)
    ->through([
        function($comment, $next) {
            $comment = trim($comment);
            return $next($comment);
        },
        function($comment, $next) {
            $comment = str_replace(':-)', '😁', $comment);
            return $next($comment);
        },
    ])
    ->thenReturn();

?><?php # real world example: Build the Query

class ItemsController extends Controller
{
    public function __invoke(Request $request)
    {
        $itemsQuery = Item::query();

        if ($request->has('search')) {
            $itemQuery->whereLike('name', "%{$request->get('search')}%");
        }

        if ($request->has('category')) {
            $itemQuery->where('category', $request->get('category'));
        }

        if ($request->has('best_deal')) {
            $itemQuery->where('best_deal', true);
        }

        return view('items', ['items' => $itemQuery->get()]);

       #return view('items', ['items' => Item::all()]);
    }
}

?><?php # Use 'when' to Build the Query

class ItemsController extends Controller
{
    public function __invoke(Request $request)
    {
        $itemsQuery = Item::query()
            ->when($request->search, function($query, $search) {
                $query->whereLike('name', "%$search%");
            })
            ->when($request->category, function($query, $category) {
                $query->where('category', $category);
            })
            ->when($request->best_deal, function($query, $bestDeal) {
                $query->where('best_deal', true);
            })

        return view('items', ['items' => $itemQuery->get()]);
    }
}

?><?php # use Pipeline

class ItemsController extends Controller
{
    public function __invoke(Request $request)
    {
        $pipes = [
            new App\Filters\SearchFilter($request->get('search')),
            new App\Filters\CategoryFilter($request->get('catetory')),
            new App\Filters\BestDealFilter($request->get('best_deal')),
        ];

        $items = Pipeline::send(Item::query())
            ->through($pipes)
            ->thenReturn()
            ->get();

        return view('items', ['items' => $items]);
    }
}

namespace App\Filters;

class SearchFilter
{
    public function __construct(private ?string $value) {}

    public function __invoke(Builder $query, $next)
    {
        if (! $this->value) {
            return $next($query);
        }

        $query->whereLike('name', "%{$this->value}%");

        return $next($query);
    }
}

class CategoryFilter
{
    public function __construct(private ?string $value) {}

    public function __invoke(Builder $query, $next)
    {
        if (! $this->value) {
            return $next($query);
        }

        $query->where('category', $this->value);

        return $next($query);
    }
}

class BestDealFilter
{
    public function __construct(private ?string $value) { }

    public function __invoke(Builder $query, $next)
    {
        if (! $this->value) {
            return $next($query);
        }

        $query->where('best_deal', true);

        return $next($query);
    }
}
?>
