<template>
  <button type="button" @click="$emit('click')"><slot></slot></button>
</template>

<script>
export default {
  name: 'Button',
};
</script>
