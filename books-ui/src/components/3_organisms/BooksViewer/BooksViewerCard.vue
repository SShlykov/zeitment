<script>
import moment from 'moment';

export default {
  name: 'BooksViewerCard',
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
    // is_public: {
    //   type: Boolean,
    //   required: true
    // },
    // publication: {
    //   type: String,
    //   required: false
    // },
    // image_link: {
    //   type: String,
    //   required: false
    // },
    // map_link: {
    //   type: String,
    //   required: false
    // },
    // map_params_id: {
    //   type: String,
    //   required: false
    // },
    // variables: {
    //   type: Array,
    //   required: false
    // },
    // updated_at: {
    //   type: String,
    //   required: true
    // }
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
      <div class="h-full w-full p-4">
        <div>
          <h3
            v-tippy="{ content: title }"
            test-id="title"
            class="text-lg font-bold mb-1 truncate transition-all group-hover:text-slate-700"
          >
            {{ title }}
          </h3>
          <p
            test-id="lastUpdated"
            class="text-xs group-hover:text-slate-600 transition-all"
          >
            Обновлено {{ lastUpdated }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
