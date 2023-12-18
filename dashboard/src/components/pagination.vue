<!-- Pagination.vue -->
<template>
    <div class="pagination">
        <button @click="changePage('prev')" :disabled="currentPage === 1">Previous</button>
        <span>{{ currentPage }} </span>
        <button @click="changePage('next')" :disabled="currentPage === totalPages">Next</button>
    </div>
</template>

<script>
export default {
    name: "PagView",
    props: {
        perPage: {
            type: Number,
            required: true,
        },
        currentPage: {
            type: Number,
            required: true,
        },
    },

    methods: {
        changePage(action) {
            let newPage;
            if (action === 'prev') {
                newPage = this.currentPage - 1;

            } else if (action === 'next') {
                newPage = this.currentPage + 1;

            }

            if (newPage > 0) {
                this.$emit('page-change', newPage); // Phát ra sự kiện với trang mới
                console.log('Current Page:', newPage);
            }
        },
    },
};
</script>
  
<style scoped>
.pagination {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-top: 20px;
}

.pagination button {
    padding: 8px 12px;
    cursor: pointer;
}

.pagination button:disabled {
    cursor: not-allowed;
    opacity: 0.5;
}
</style>
  