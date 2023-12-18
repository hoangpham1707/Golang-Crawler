<!-- ListSource.vue -->
<template>
    <div class="container mx-auto p-4">
        <h1 class="text-2xl font-semibold mb-4">List mention</h1>
        <!-- <h3 class="text-2xl font-semibold mb-4">Total Web: {{ totalWeb }} Total Web: {{ totalWeb }} Total Facebook: {{ totalFb }}</h3> -->
        <ul class="entry-list">
            <li v-for="entry in entries" :key="entry._id" class="mb-2">
                <div class="flex items-center justify-between">
                    <div class="entry-container">
                        <p>{{ entry.content }}</p>
                    </div>
                </div>
            </li>
        </ul>
        <!-- Nút phân trang -->
        <div class="mt-4 flex items-center justify-center">
            <button @click="changePage(-1)" :disabled="pagination.page === 1">Trở lại</button>
            <span class="mx-2">Trang {{ pagination.page }} / {{ totalPages }}</span>
            <button @click="changePage(1)" :disabled="pagination.page === totalPages">Tiếp</button>
        </div>

    </div>
</template>

  
<script>
import axios from 'axios';

export default {
    name: 'HomeView',
    data() {
        return {
            entries: [],
            // totalWeb: 0,
            // totalFb: 0,
            pagination: {
                page: 1,
                pageSize: 10,
            },
        };
    },
    computed: {
        totalPages() {
            return Math.ceil(this.pagination.total / this.pagination.pageSize);
        },
    },
    created() {
        this.fetchMentions();
        // this.fetchMentionsCount();
    },
    methods: {

        async fetchMentions() {
            try {
                const response = await axios.request({
                    method: 'GET',
                    url: `http://localhost:6006/api/v1/common/mentions`,
                    params: {
                        page: this.pagination.page,
                        pageSize: this.pagination.pageSize,
                    },
                });
                console.log({ response });
                this.entries = response.data.data.entries;
                this.pagination = response.data.data.pagination;
                // console.log('Current Page:', this.pagination.page);
            } catch (error) {
                console.error('Error fetching sources:', error);
            }
        },
        // async fetchMentionsCount() {
        //     try {
        //         const response = await axios.request({
        //             method: 'GET',
        //             url: `http://localhost:6006/api/v1/common/mentions/count`
        //         });
        //         console.log({ response });
        //         this.totalWeb = response.data.data.totalWeb;
        //         this.totalFb = response.data.data.totalFacebook;
        //         console.log(response.data.data.totalFacebook);
        //     } catch (error) {
        //         console.error('Error fetching sources:', error);
        //     }
        // },
        changePage(offset) {
            const newPage = this.pagination.page + offset;
            if (newPage > 0 && newPage <= this.totalPages) {
                this.pagination.page = newPage;
                this.fetchMentions();
            }
        },
    },
};
</script>
  
<style scoped>
/* Add styling for the entry container */
.entry-container {
    max-height: 110px;
    /* Adjust the height as needed */
    overflow: hidden;
    border: 1px solid #ddd;
    padding: 10px;
    margin-bottom: 10px;
}

.entry-list {
    display: flex;
    flex-wrap: wrap;
    /* Optional: Wrap to the next line if the items don't fit */
    justify-content: space-between;
    /* Optional: Adjust the alignment */
    list-style: none;
    /* Optional: Remove default list styling */
    padding: 0;
    /* Optional: Remove default padding */
}

.pagination {
    display: flex;
    justify-content: space-between;
}
</style>
