Let's consider the following "counter" component written in VueJs:

<script>
    export default {
        data: {
            count: 0
        },
        methods: {
            increment() {
                this.count++
            },
            decrement() {
                this.count--
            },
        },
    }
</script>

<template>
    <div>
        <button @click="increment">+</button>
        <button @click="decrement">-</button>

        <span>{{ count }}</span>
    </div>
</template> 

Now, let's see how we would accomplish the exact same thing with a Livewire component.

app/Http/Livewire/Counter.php

<?php

use Livewire\Component;

class Counter extends Component
{
    public $count = 0;

    public function increment()
    {
        $this->count++;
    }

    public function decrement()
    {
        $this->count--;
    }

    public function render()
    {
        return view('livewire.counter');
    }
}
?>

resources/views/livewire/counter.blade.php

<div>
    <button wire:click="increment">+</button>
    <button wire:click="decrement">-</button>

    <span>{{ $this->count }}</span>
</div> 

resources/views/index.blade.php

    @livewire('counter')
