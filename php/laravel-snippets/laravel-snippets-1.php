<?php

Laravel Snippets & Methods

# 1. Determining if the record on firstOrCreate was new or not

$product = Product::firstOrCreate(...)

if ($product->wasRecentlyCreated()) {
    // New product
} else {
    // Existing product
}

# 2. Find related IDs on a BelongsToMany Relationship

$user->roles()->allRelatedIds()->toArray();

# 3. abort_unless()

// Instead of
public function show($item) {
    if (// User can not do this thing) {
        return false;
    }
    
    // Else do this
}

// Do this
public function show($item) {
    abort_unless(Gate::allows('do-thing', $item), 403);
    
    // Actual logic
}

// Another use case, make sure the user is logged in
abort_unless(Auth::check(), 403);

# 4. Model Keys

User::all()->pluck('id')->toArray();

// In most cases, however, this can be shortened. Like this:

User::all()->modelKeys();

# 5. throw_if()

throw_if(
    !Hash::check($data['current_password'], $user->password),
    new Exception(__('That is not your old password.'))
);

# 6. Dump all columns of a table

Schema::getColumnListing('table')

# 7. Redirect to external domains

return redirect()->away('https://www.google.com');

# 8. Request exists() vs has()

// http://example.com?popular

$request->exists('popular') // true
$request->has('popular') // false

http://example.com?popular=foo

$request->exists('popular') // true
$request->has('popular') // true

# 9. @isset

// From
@if (isset($records))
    // $records is defined and is not null
@endif

// To
@isset($records)
    // $records is defined and is not null
@endisset

# 10. @empty

// From
@if (empty($records))
    // $records is "empty"
@endif

// To
@empty($records)
    // $records is "empty"
@endempty

# 11. @forelse

// From
@if ($users->count())
    @foreach ($users as $user)
    
    @endforeach
@else

@endif

// To
@forelse ($users as $user)
    {{ $user->name }}
@empty
    <p>No Users</p>
@endforelse

# 12. array_wrap()

$posts = is_array($posts) ? $posts : [$posts];

// Same as
$posts = array_wrap($posts);

// Inline
foreach (array_wrap($posts) as $post) {
    // ..
}

# 13. optional()

The optional() helper allows you to access properties or call methods on an object. If the given object is null, properties and methods will return null instead of causing an error.

// User 1 exists, with account
$user1 = User::find(1);
$accountId = $user1->account->id; // 123

// User 2 exists, without account
$user2 = User::find(2);
$accountId = $user2->account->id; // PHP Error: Trying to get property of non-object

// Fix without optional()
$accountId = $user2->account ? $user2->account->id : null; // null
$accountId = $user2->account->id ?? null; // null

// Fix with optional()
$accountId = optional($user2->account)->id; // null

# 14. data_get()

The data_get() helper allows you to get a value from an array or object with dot notation. This functions similarly to array_get() as well. The optional third parameter can be used to supply a default value if the key is not found.

$array = ['albums' => ['rock' => ['count' => 75]]];

$count = data_get($array, 'albums.rock.count'); // 75
$avgCost = data_get($array, 'albums.rock.avg_cost', 0); // 0

$object->albums->rock->count = 75;

$count = data_get($object, 'albums.rock.count'); // 75
$avgCost = data_get($object, 'albums.rock.avg_cost', 0); // 0

# 15. push()

You can save a model and its corresponding relationships using the push() method.

$user = User::first();
$user->name = "Peter";

$user->phone->number = '1234567890';

$user->push(); // This will update both user and phone record in DB
