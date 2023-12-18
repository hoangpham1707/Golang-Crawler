
<template>
    <div class="container mx-auto p-1">
        <p class="text-2xl font-semibold mb-1">Khung HTML</p>
        <table class="w-full border">
            <thead>
                <tr class="bg-gray-200">
                    <th class="border px-4 py-2">URL Start</th>
                    <th class="border px-4 py-2">Box Element</th>
                    <th class="border px-4 py-2">Title Element</th>
                    <th class="border px-4 py-2">Link Element</th>
                    <th class="border px-4 py-2">Time Element</th>
                    <th class="border px-4 py-2">Description Element</th>
                    <th class="border px-4 py-2">Category Element</th>
                    <th class="border px-4 py-2">Content Element</th>
                    <th class="border px-4 py-2">Check Time</th>
                    <th class="border px-4 py-2">Check Description</th>
                    <th class="border px-4 py-2">Check Category</th>
                    <th class="border px-4 py-2">Hành động</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="entry in entries" :key="entry.id" class="border-b">
                    <td class="border px-4 py-2">{{ entry.url_start }}</td>
                    <td class="border px-4 py-2">{{ entry.box_element }}</td>
                    <td class="border px-4 py-2">{{ entry.title_element }}</td>
                    <td class="border px-4 py-2">{{ entry.link_element }}</td>
                    <td class="border px-4 py-2">{{ entry.time_element }}</td>
                    <td class="border px-4 py-2">{{ entry.description_element }}</td>
                    <td class="border px-4 py-2">{{ entry.category_element }}</td>
                    <td class="border px-4 py-2">{{ entry.content_element }}</td>
                    <td class="border px-4 py-2">{{ entry.check_time }}</td>
                    <td class="border px-4 py-2">{{ entry.check_desc }}</td>
                    <td class="border px-4 py-2">{{ entry.check_category }}</td>
                    <td class="border px-4 py-2">
                        <button @click="editUrlNext(entry.id)" class="bg-blue-500 text-white px-2 py-1 rounded">Sửa</button>
                        <button @click="deleteUrlNext(entry.id)"
                            class="ml-2 bg-red-500 text-white px-2 py-1 rounded">Xóa</button>
                    </td>
                </tr>
            </tbody>
        </table>
        <form @submit.prevent="addKeyword" class="mt-4 flex items-center">
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
    name: 'ListLink',

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
    methods: {
        async fetchSources() {
            try {
                const response = await axios.request({
                    methods: 'GET',
                    url: `http://localhost:6006/api/v1/common/linkNexts`,
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
            this.$router.push({ path: `/createLinkNext` });
        },
        editUrlNext(entryId) {
            this.$router.push({ path: `/detailLinkNext/${entryId}` });
        },
        async deleteUrlNext(id) {
            try {
                await axios.delete(`http://localhost:6006/api/v1/common/linkNexts/${id}`);
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