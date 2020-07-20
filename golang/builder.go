// The builder pattern makes it really hard to return errors during a build
// step. Consider using functional options instead, which work much better with
// incremental errors.

// Let's say you have the builder pattern.
type T struct { ... }

func (t *T) A(arg string) *T {
  // do stuff here with t and arg
  return t
}

// You use it with somethign like:
t := &T{}
t = t.A("hello").B(...).C(...)

// Moving to functional options is pretty easy.

// Step 1: Replace all methods with functions that return a closure of the format func(*T) *T
func A(arg string) func(*T) *T {
  return func(t *T) *T {
    // do stuff here with t and arg
    return t
  }
}

// Step 2: Create a function to build.
func New(fns ...func(*T) *T) *T {
  t := &T{}
  for _, fn := range fns {
    t = fn(t)
  }
  return t
}

// Step 3: Use your functional options.
t := New(A("hello"), B(...), C(...))


// As stated earlier, this pattern is much easier to return individual
// build-step errors.
func A(arg string) func(*T) (*T, error) {
  return func(*T) (*T, error) {
    if arg == "" { return nil, fmt.Errorf("empty arg") }
    // do stuff here with t and arg
    return t, nil
  }
}

func New(fns ...func(*T) (*T, error)) (*T, error) {
  t := &T{}
  var err error
  for _, fn := range fns {
    t, err = fn(t)
    if err != nil {
        return nil, err
    }
  }
  return t, nil
}

// No need to deal with errors on each individual build step.
t, err := New(DoSomething("hello"), AnotherThing(...))