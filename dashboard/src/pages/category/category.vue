<template>
    <div class="container mx-auto p-4">
        <h1 class="text-2xl font-semibold mb-4">Danh sách danh mục</h1>
        <table class="w-full border">
            <thead>
                <tr class="bg-gray-200">
                    <th class="border px-4 py-2">STT</th>
                    <th class="border px-4 py-2">Tên</th>
                    <th class="border px-4 py-2">Hành động</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(entry, index) in paginatedEntries" :key="entry.id" class="border-b">
                    <td class="border px-4 py-2">{{ (currentPage - 1) * pageSize + index + 1 }}</td>
                    <td class="border px-4 py-2">{{ entry.name }}</td>

                    <td class="border px-4 py-2">
                        <button @click="editCategory(entry._id)"
                            class="bg-blue-500 text-white px-2 py-1 rounded">Sửa</button>
                        <button @click="deleteCategory(entry._id)"
                            class="ml-2 bg-red-500 text-white px-2 py-1 rounded">Xóa</button>
                    </td>
                </tr>
            </tbody>
        </table>


        <form @submit.prevent="addCategory(newCategory)" class="mt-4 flex items-center">
            <label for="newCategory" class="mr-2">Tên danh mục</label>
            <input v-model="newCategory" type="text" id="newCategory" class="border px-2 py-1" required />
            <button type="submit" class="ml-2 bg-green-500 text-white px-2 py-1 rounded">Thêm</button>
        </form>

        <!-- Nút phân trang -->
        <div class="pagination mt-4 flex items-center justify-center">
            <button @click="changePage(-1)" :disabled="currentPage === 1">Trở lại</button>
            <span class="mx-2">Trang {{ currentPage }} / {{ totalPages }}</span>
            <button @click="changePage(1)" :disabled="currentPage === totalPages">tiếp</button>
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
            currentPage: 1, // Trang hiện tại
            pageSize: 10,
        };
    },
    computed: {
        paginatedEntries() {
            const startIndex = (this.currentPage - 1) * this.pageSize;
            const endIndex = startIndex + this.pageSize;
            return this.entries.slice(startIndex, endIndex);
        },
        totalPages() {
            return Math.ceil(this.entries.length / this.pageSize);
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
                    url: `http://localhost:6006/api/v1/common/categories`,
                });
                console.log({ response })
                this.entries = response.data.data.entries;
            } catch (error) {
                console.error('Error fetching sources:', error);
            }
        },
        async addCategory(newCategory) {
            if (newCategory) {
                try {
                    await axios.post('http://localhost:6006/api/v1/common/categories', { name: newCategory });
                    console.log("Name" + newCategory)
                    newCategory = '';
                    await this.fetchSources();
                    window.location.reload();
                } catch (error) {
                    console.error('Error adding source:', error);
                }
            }
        },
        editCategory(entryId) {
            this.$router.push({ path: `/detailCategory/${entryId}` });
        },
        async deleteCategory(id) {
            try {
                await axios.delete(`http://localhost:6006/api/v1/common/categories/${id}`);
                await this.fetchSources();
            } catch (error) {
                console.error('Error deleting source:', error);
            }
        },
        changePage(offset) {
            const newPage = this.currentPage + offset;
            if (newPage > 0 && newPage <= this.totalPages) {
                this.currentPage = newPage;
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