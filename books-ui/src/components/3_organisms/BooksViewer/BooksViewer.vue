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
        :isPublic="item.isPublic"
        :publication="item.publication"
        :imageLink="item.imageLink"
        :mapLink="item.mapLink"
        :mapParamsId="item.mapParamsId"
        :variables="item.variables"
        :updatedAt="item.updatedAt"
      />
    </BooksViewerContainer>
  </div>
</template>
