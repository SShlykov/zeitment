<script>
import moment from 'moment';
import { directive } from 'vue-tippy';

export default {
  name: 'BooksViewerCard',
  directives: {
    tippy: directive
  },
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
      default: ""
    },
    description: {
      type: String,
      default: ""
    },
    updatedAt: {
      type: String,
      required: true
    },
    bookLink: {
      type: String,
      required: true
    }
  },
  computed: {
    formattedUpdatedAt() {
      return moment(this.updatedAt).format('MMMM Do YYYY, h:mm:ss a');
    },
    lastUpdated() {
      moment.locale("ru");
      return moment(this.updatedAt).fromNow();
    }
  },
  mounted() {},
  methods: {}
}

</script>

<template>
  <a
    :href="bookLink"
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
  </a>
</template>
