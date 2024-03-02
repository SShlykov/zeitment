<script>
export default {
  name: 'BackgroundImage',
  props: {
    image_link: {
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
      <img class="object-cover w-full h-full" :src="image_link" :alt="alt">
    </figure>
  </div>
</template>
