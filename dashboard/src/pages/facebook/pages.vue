<!-- ListSource.vue -->
<template>
    <div class="container mx-auto p-4">
        <h1 class="text-2xl font-semibold mb-4">List Facebook</h1>
        <h1 class="text-2xl font-semibold mb-4">Số lượng còn lại: {{ pagination.total }}</h1>

        <ul class="entry-list">
            <li v-for="entry in entries" :key="entry._id" class="mb-2">
                <div class="flex items-center justify-between">
                    <div class="entry-container">
                        <p>{{ entry._id }}</p>
                        <p>{{ entry.content }}</p>

                    </div>
                    <div class="flex items-center">
                        <select v-model="entry.selectedLabel">
                            <option v-for="label in labels" :key="label.slug" :value="label.slug">{{ label.name }}</option>
                        </select>
                        <button @click="assignLabel(entry._id, entry.selectedLabel, currentPage)"
                            class="ml-2 bg-green-500 text-white px-2 py-1 rounded">Gán Nhãn</button>
                        <button @click="readMention(entry._id)"
                            class="ml-2 bg-blue-500 text-white px-2 py-1 rounded">Đọc</button>
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
    name: 'ListAddLabel',
    data() {
        return {
            entries: [],
            labels: [], // Thêm mảng này để lưu trữ các nhãn
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
        this.fetchLabels();
    },
    watch: {
        '$route.query.account_id': {
            handler: 'fetchMentions', // Gọi phương thức fetchMentions khi có thay đổi
            immediate: true, // Để nó được gọi ngay từ lúc ban đầu
        },
    },
    methods: {
        async fetchLabels() {
            try {
                const response = await axios.get('http://localhost:6006/api/v1/common/labels');
                this.labels = response.data.data.entries;
                this.entries.forEach(entry => {
                    entry.selectedLabel = this.labels[0]?.name || ''; // Default to the first label
                });
            } catch (error) {
                console.error('Lỗi khi lấy danh sách nhãn:', error);
            }
        },
        async assignLabel(entryId, selectedLabel, currentPage) {
            try {
                await axios.put(`http://localhost:6006/api/v1/common/mentions/${entryId}/assign-label?label_name=${selectedLabel}`);
                // Lưu giá trị vào localStorage để giữ nguyên sau khi reload

                await this.fetchSources(currentPage);
            } catch (error) {
                console.error('Lỗi khi gán nhãn:', error);
            }
        },

        async fetchMentions() {
            try {
                const account_id = this.$route.query.account_id || '';
                //this.accountID = '656032122d8638c72626a67e'
                const response = await axios.request({
                    method: 'GET',
                    url: `http://localhost:6006/api/v1/common/mentions/addLabelFace`,
                    params: {
                        account_id,
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
        async readMention(entryId) {
            const account_id = this.$route.query.account_id || '';
            this.$router.push({
                path: `/sentences/${entryId}`,
                query: { account_id }
            });
        },
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
