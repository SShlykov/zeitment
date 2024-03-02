<script>
import {mapActions, mapGetters} from "vuex";
import BooksViewerCard from './BooksViewerCard.vue'
import BooksViewerContainer from './BooksViewerContainer.vue'

export default {
  name: 'BooksViewer',
  components: {BooksViewerCard, BooksViewerContainer},
  data() {},
  computed: {
    ...mapGetters('userBooks', ['booksList'])
  },
  mounted() {
    this.fetchUserBooks()
  },
  methods: {
    ...mapActions('userBooks', ['fetchUserBooks'])
  }
}

</script>

<template>
  <div class="">
    <BooksViewerContainer>
      <BooksViewerCard 
        v-for="item in booksList"
        :id="item.id"
        :key="item.id"
        :owner="item.owner"
        :title="item.title"
        :author="item.author"
        :description="item.description"
        :is_public="item.is_public"
        :publication="item.publication"
        :image_link="item.image_link"
        :map_link="item.map_link"
        :map_params_id="item.map_params_id"
        :variables="item.variables"
        :updated_at="item.updated_at"
      />
    </BooksViewerContainer>
  </div>
</template>
