// golang-优秀代码赏析-eventually等待获取最后存储起来的value

// 并发获取最后存储起来的value，如果获取不到就会陷入等待一段时间。

// Package eventual provides values that eventually have a value.
package eventual

import (
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// Value is an eventual value, meaning that callers wishing to access the value
// block until the value is available.
type Value interface {
	// Set sets this Value to the given val.
	Set(val interface{})

	// Get waits up to timeout for the value to be set and returns it, or returns
	// nil if it times out or Cancel() is called. valid will be false in latter
	// case. If timeout is 0, Get won't wait. If timeout is -1, Get will wait
	// forever.
	Get(timeout time.Duration) (ret interface{}, valid bool)

	// Cancel cancels this value, signaling any waiting calls to Get() that no
	// value is coming. If no value was set before Cancel() was called, all future
	// calls to Get() will return nil, false. Subsequent calls to Set after Cancel
	// have no effect.
	Cancel()
}

// Getter is a functional interface for the Value.Get function
type Getter func(time.Duration) (interface{}, bool)

type value struct {
	state   atomic.Value
	waiters []chan interface{}
	mutex   sync.Mutex
}

type stateholder struct {
	val      interface{}
	set      bool
	canceled bool
}

// NewValue creates a new Value.
func NewValue() Value {
	result := &value{waiters: make([]chan interface{}, 0)}
	result.state.Store(&stateholder{})
	return result
}

// DefaultGetter builds a Getter that always returns the supplied value.
func DefaultGetter(val interface{}) Getter {
	return func(time.Duration) (interface{}, bool) {
		return val, true
	}
}

// DefaultUnsetGetter builds a Getter that always !ok.
func DefaultUnsetGetter() Getter {
	return func(time.Duration) (interface{}, bool) {
		return nil, false
	}
}

func (v *value) Set(val interface{}) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	state := v.getState()
	settable := !state.canceled
	if settable {
		v.setState(&stateholder{
			val:      val,
			set:      true,
			canceled: false,
		})

		if v.waiters != nil {
			// Notify anyone waiting for value
			for _, waiter := range v.waiters {
				waiter <- val
			}
			// Clear waiters
			v.waiters = nil
		}
	}
}

func (v *value) Cancel() {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	state := v.getState()
	v.setState(&stateholder{
		val:      state.val,
		set:      state.set,
		canceled: true,
	})

	if v.waiters != nil {
		// Notify anyone waiting for value
		for _, waiter := range v.waiters {
			close(waiter)
		}
		// Clear waiters
		v.waiters = nil
	}
}

func (v *value) Get(timeout time.Duration) (ret interface{}, valid bool) {
	state := v.getState()

	// First check for existing value using atomic operations (for speed)
	if state.set {
		// Value found, use it
		return state.val, true
	} else if state.canceled {
		// Value was canceled, return false
		return nil, false
	}

	if timeout == 0 {
		// Don't wait
		return nil, false
	}

	// If we didn't find an existing value, try again but this time using locking
	v.mutex.Lock()
	state = v.getState()

	if state.set {
		// Value found, use it
		v.mutex.Unlock()
		return state.val, true
	} else if state.canceled {
		// Value was canceled, return false
		v.mutex.Unlock()
		return nil, false
	}

	if timeout == -1 {
		// Wait essentially forever
		timeout = time.Duration(math.MaxInt64)
	}

	// Value not found, register to be notified once value is set
	valCh := make(chan interface{}, 1)
	v.waiters = append(v.waiters, valCh)
	v.mutex.Unlock()

	// Wait up to timeout for value to get set
	select {
	case v, ok := <-valCh:
		return v, ok
	case <-time.After(timeout):
		return nil, false
	}
}

func (v *value) getState() *stateholder {
	state := v.state.Load()
	if state == nil {
		return nil
	}
	return state.(*stateholder)
}

func (v *value) setState(state *stateholder) {
	v.state.Store(state)
}

// test

package eventual

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/getlantern/grtrack"
	"github.com/stretchr/testify/assert"
)

const (
	concurrency = 200
)

func TestSingle(t *testing.T) {
	goroutines := grtrack.Start()
	v := NewValue()
	go func() {
		time.Sleep(20 * time.Millisecond)
		v.Set("hi")
	}()

	r, ok := v.Get(0)
	assert.False(t, ok, "Get with no timeout should have failed")

	r, ok = v.Get(10 * time.Millisecond)
	assert.False(t, ok, "Get with short timeout should have timed out")

	r, ok = v.Get(-1)
	assert.True(t, ok, "Get with really long timeout should have succeeded")
	assert.Equal(t, "hi", r, "Wrong result")

	// Set a different value
	v.Set("bye")
	r, ok = v.Get(0)
	assert.True(t, ok, "Subsequent get with no timeout should have succeeded")
	assert.Equal(t, "bye", r, "Value should have changed")

	goroutines.CheckAfter(t, 50*time.Millisecond)
}

func TestNoSet(t *testing.T) {
	goroutines := grtrack.Start()
	v := NewValue()

	_, ok := v.Get(10 * time.Millisecond)
	assert.False(t, ok, "Get before setting value should not be okay")

	goroutines.CheckAfter(t, 50*time.Millisecond)
}

func TestCancelImmediate(t *testing.T) {
	v := NewValue()
	go func() {
		time.Sleep(10 * time.Millisecond)
		v.Cancel()
	}()

	_, ok := v.Get(200 * time.Millisecond)
	assert.False(t, ok, "Get after cancel should have failed")
}

func TestCancelAfterSet(t *testing.T) {
	v := NewValue()
	v.Set(5)
	r, ok := v.Get(10 * time.Millisecond)
	assert.True(t, ok, "Get before cancel should have succeeded")
	assert.Equal(t, 5, r, "Get got wrong value before cancel")

	v.Cancel()
	r, ok = v.Get(0)
	assert.True(t, ok, "Get after cancel should have succeeded")
	assert.Equal(t, 5, r, "Get got wrong value after cancel")

	v.Set(10)
	r, _ = v.Get(0)
	assert.Equal(t, 5, r, "Set after cancel should have no effect")
}

func BenchmarkGet(b *testing.B) {
	v := NewValue()
	go func() {
		time.Sleep(20 * time.Millisecond)
		v.Set("hi")
	}()

	for i := 0; i < b.N; i++ {
		v.Get(20 * time.Millisecond)
	}
}

func TestConcurrent(t *testing.T) {
	goroutines := grtrack.Start()
	v := NewValue()

	var sets int32

	go func() {
		var wg sync.WaitGroup
		wg.Add(1)
		// Do some concurrent setting to make sure that it works
		for i := 0; i < concurrency; i++ {
			go func() {
				// Wait for waitGroup so that all goroutines run at basically the same
				// time.
				wg.Wait()
				v.Set("hi")
				atomic.AddInt32(&sets, 1)
			}()
		}
		wg.Done()
	}()

	for i := 0; i < concurrency; i++ {
		go func() {
			r, ok := v.Get(200 * time.Millisecond)
			assert.True(t, ok, "Get should have succeed")
			assert.Equal(t, "hi", r, "Wrong result")
		}()
	}

	goroutines.CheckAfter(t, 50*time.Millisecond)
	assert.EqualValues(t, concurrency, atomic.LoadInt32(&sets), "Wrong number of successful Sets")
}