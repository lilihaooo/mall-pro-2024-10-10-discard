<script setup>
import { useRoute } from 'vue-router';

const route = useRoute();

const id = route.query.id
</script>

<template>

  {{id}}

</template>

<style scoped lang="scss">

</style>