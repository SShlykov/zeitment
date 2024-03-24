<script>
import { VueDraggableNext } from 'vue-draggable-next'

export default {
  name: 'BooksCompendium',
  components: {
    draggable: VueDraggableNext,
  },
  data() {
    return {
      enabled: true,
      list: [
        { title: 'John',   id: 1, items: [{
          title: "John-Item 1",
        },
        {
          title: "John-Item 2",
        },
        {
          title: "John-Item 3",
        }
        ] },
        { title: 'Joao',   id: 2, items: [{
          title: "Joao-Item 1",
        },
        {
          title: "Joao-Item 2",
        },
        {
          title: "Joao-Item 3",
        }
        ] },
        { title: 'Jean',   id: 3, items: [{
          title: "Jean-Item 1",
        }, {
          title: "Jean-Item 2",
        },
        {
          title: "Jean-Item 3",
        }
        ] },
        { title: 'Gerard', id: 4, items: [
          {
            title: "Gerard-Item 1",
          },
          {
            title: "Gerard-Item 2",
          },
          {
            title: "Gerard-Item 3",
          }
        ] },
      ],
      dragging: false,
    }
  },
  computed: {},
  mounted() {},
  methods: {
    log(event) {
      console.log(event)
    },
  },
}

</script>

<template>
  <div class="">
    <draggable
      class="w-full"
      :list="list"
      animation="'150'"
      @change="log"
    >
      <div
        v-for="element in list"
        :key="element.title"
        class="bg-gray-300 m-1 p-3 rounded-md text-center"
      >
        <span>
          {{ element.title }}
        </span>
        <draggable
          class="list-group"
          :list="element.items"
          animation="'150'"
          @change="log"
          group="nested"
        >
          <div
            v-for="item in element.items"
            :key="item.title"
            class="list-group-item bg-gray-200 m-1 p-3 rounded-md text-center"
          >
            <span>
              {{ item.title }}
            </span>
          </div>
        </draggable>
      </div>
    </draggable>
  </div>
</template>
