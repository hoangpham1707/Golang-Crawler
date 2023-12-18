
<template>
    <div class="container mx-auto p-1">
        <p class="text-2xl font-semibold mb-1">Danh sách từ khóa Crawl</p>
        <table class="w-full border">
            <thead>
                <tr class="bg-gray-200">
                    <th class="border px-4 py-2">Tên</th>
                    <th class="border px-4 py-2">Hành động</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="entry in entries" :key="entry.id" class="border-b">
                    <td class="border px-4 py-2">{{ entry.name }}</td>
                    <td class="border px-4 py-2">
                        <button @click="editKeyword(entry.id)"
                            class="ml-2 bg-blue-500 text-white px-2 py-1 rounded">Sửa</button>
                        <button @click="deleteKeyword(entry.id)"
                            class="ml-2 bg-red-500 text-white px-2 py-1 rounded">Xóa</button>
                    </td>
                </tr>
            </tbody>
        </table>


        <form @submit.prevent="addKeyword" class="mt-4 flex items-center">
            <label for="newKeyword" class="mr-2">Từ khóa</label>
            <input v-model="newKeyword" type="text" id="newKeyword" class="border px-2 py-1" required />
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
    name: 'ListKeyword',

    data() {
        return {
            entries: [],
            newSource: '',
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

        this.fetchSources();
    },
    methods: {
        async fetchSources() {
            try {
                const response = await axios.request({
                    methods: 'GET',
                    url: `http://localhost:6006/api/v1/common/keywords`,
                    params: {
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
        async addKeyword() {
            if (this.newKeyword) {
                try {
                    await axios.post('http://localhost:6006/api/v1/common/keywords', { name: this.newKeyword });

                    this.newKeyword = '';
                    await this.fetchSources();
                    window.location.reload();


                } catch (error) {
                    console.error('Error adding source:', error);
                }
            }
        },
        editKeyword(entryId) {
            this.$router.push({ path: `/detailKeyword/${entryId}` });
        },
        async deleteKeyword(id) {
            try {
                await axios.delete(`http://localhost:6006/api/v1/common/keywords/${id}`);
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
    },
};
</script>
  
<style>
.pagination {
    display: flex;
    justify-content: space-between;
}
</style>