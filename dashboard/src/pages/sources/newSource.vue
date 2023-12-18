<template>
    <div class="container mx-auto p-1">
        <p class="text-2xl font-semibold mb-1">Danh sách nguồn</p>
        <table class="w-full border">
            <thead>
                <tr class="bg-gray-200">
                    <th class="border px-4 py-2">STT</th>
                    <th class="border px-4 py-2">Ảnh</th>
                    <th class="border px-4 py-2">Đừng dẫn</th>
                    <th class="border px-4 py-2">Số lượng bài</th>
                    <th class="border px-4 py-2">Hành động</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(entry, index) in entries" :key="entry.id" class="border-b">
                    <td class="border px-4 py-2">{{ (pagination.page - 1) * pagination.pageSize + index + 1 }}</td>
                    <img :src="entry.avatar" alt="Source Avatar" class="mr-2 w-8 h-8 rounded-full source-avatar">
                    <td class="border px-4 py-2">{{ entry.url_start }}</td>
                    <td class="border px-4 py-2">{{ entry.crawl_count }}</td>
                    <td class="border px-4 py-2">
                        <button @click="editSource(entry.id)"
                            class="ml-2 bg-blue-500 text-white px-2 py-1 rounded">Sửa</button>
                        <button @click="deleteSource(entry.id)"
                            class="ml-2 bg-red-500 text-white px-2 py-1 rounded">Xóa</button>
                    </td>
                </tr>
            </tbody>
        </table>

        <form @submit.prevent="addSource" class="mt-4 flex items-center">
            <button type="submit" class="ml-2 bg-green-500 text-white px-2 py-1 rounded">Thêm</button>
        </form>
        <div class="pagination mt-4 flex items-center justify-center">
            <button @click="changePage(-1)" :disabled="pagination.page === 1">Trở lại</button>
            <span class="mx-2">Trang {{ pagination.page }} / {{ totalPages }}</span>
            <button @click="changePage(1)" :disabled="pagination.page === totalPages">Tiếp</button>
        </div>
    </div>
</template>



<script>
import axios from 'axios';

export default {
    name: 'ListSource',

    data() {
        return {
            entries: [],
            newSource: '',
            pagination: {
                page: 1,
                pageSize: 5,

            },
        };
    },
    computed: {
        totalPages() {
            return Math.ceil(this.pagination.total / this.pagination.pageSize);
        },
    },
    created() {

        this.fetchSources();
    },
    watch: {
        '$route.query.typeSource': {
            handler: 'typeSourceChanged', // Gọi phương thức fetchMentions khi có thay đổi
            immediate: true, // Để nó được gọi ngay từ lúc ban đầu
        },
    },
    methods: {
        async typeSourceChanged() {
            // Đặt trang về 1 và tải lại dữ liệu
            this.pagination.page = 1;
            this.fetchSources();
        },
        async fetchSources() {
            try {
                const typeSource = this.$route.query.typeSource || '';
                const response = await axios.request({
                    methods: 'GET',
                    url: 'http://localhost:6006/api/v1/common/sources/type',
                    params: {
                        typeSource,
                        page: this.pagination.page,
                        pageSize: this.pagination.pageSize,
                    },
                });
                console.log({ response })
                this.entries = response.data.data.entries;
                this.pagination = response.data.data.pagination;
            } catch (error) {
                console.error('Error fetching sources:', error);
            }
        },
        async addSource() {
            const typeSource = this.$route.query.typeSource || '';
            this.$router.push({
                path: `/createSource`,
                query: { typeSource }
            });
        },
        editSource(entryId) {
            const typeSource = this.$route.query.typeSource || '';
            this.$router.push({
                path: `/detailSource/${entryId}`,
                query: { typeSource }
            });
        },
        async deleteSource(sourceId) {
            try {
                await axios.delete(`http://localhost:6006/api/v1/common/sources/${sourceId}`);
                await this.fetchSources();
            } catch (error) {
                console.error('Error deleting source:', error);
            }
        },
        changePage(offset) {
            const newPage = this.pagination.page + offset;
            if (newPage > 0 && newPage <= this.totalPages) {
                this.pagination.page = newPage;
                this.fetchSources();
            }
        },
    }
};
</script>
<style scoped>
.source-avatar {
    width: 40px;
    /* Đặt kích thước mong muốn cho avatar */
    height: 40px;
}

.pagination {
    display: flex;
    justify-content: space-between;
}
</style>
