<script>
import BackgroundImage from '@atoms/BackgroundImage.vue';
import moment from 'moment';

export default {
  name: 'BooksViewerCard',
  components: {BackgroundImage},
  props: {
    id: {
      type: String,
      required: true
    },
    owner: {
      type: String,
      required: true
    },
    title: {
      type: String,
      required: true
    },
    author: {
      type: String,
      required: true
    },
    description: {
      type: String,
      required: true
    },
    is_public: {
      type: Boolean,
      required: true
    },
    publication: {
      type: String,
      required: false
    },
    image_link: {
      type: String,
      required: false
    },
    map_link: {
      type: String,
      required: false
    },
    map_params_id: {
      type: String,
      required: false
    },
    variables: {
      type: Array,
      required: false
    },
    updated_at: {
      type: String,
      required: true
    }
  },
  data() {},
  computed: {
    formattedUpdatedAt() {
      return moment(this.updated_at).format('MMMM Do YYYY, h:mm:ss a');
    },
    lastUpdated() {
      moment.locale("ru");
      return moment(this.updated_at).fromNow();
    }
  },
  mounted() {},
  methods: {}
}

</script>

<template>
  <div
    class="relative shadow group border border-slate-300 rounded-md cursor-pointer transition-all
              hover:rounded-xl hover:border-slate-500 hover:shadow-md
               overflow-hidden
  "
  >
    <div class="card text-slate-500">
      <BackgroundImage 
        :image_link="'https://cdn.britannica.com/45/5645-050-B9EC0205/head-treasure-flower-disk-flowers-inflorescence-ray.jpg'" 
        class="opacity-0 group-hover:opacity-5 transition-all"
        :alt="title"
        v-if="false"
      />
      <div class="h-full w-full p-4">
        <div>
          <h3 class="text-lg font-bold mb-1">{{title}}</h3>
          <p test-id="lastUpdated" class="text-xs">Обновлено {{lastUpdated}}</p>
        </div>
      </div>
    </div>
  </div>
</template>
