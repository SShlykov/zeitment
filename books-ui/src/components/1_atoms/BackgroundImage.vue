<script>
export default {
  name: 'BackgroundImage',
  props: {
    imageLink: {
      type: String,
      required: true
    },
    alt: {
      type: String,
      required: true
    }
  },
}

</script>

<template>
  <div class="absolute inset-0 overflow-hidden transition-all">
    <figure class="h-full w-full overflow-hidden">
      <img
        class="object-cover w-full h-full"
        :src="imageLink"
        :alt="alt"
      >
    </figure>
  </div>
</template>
