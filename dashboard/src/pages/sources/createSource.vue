<template>
    <!-- Edit form -->
    <div class="max-w-md mx-auto mt-10 p-6 bg-white rounded shadow-md">
        <h1 class="text-2xl font-semibold mb-6">Thêm nguồn đăng tải</h1>
        <form @submit.prevent="createSource">
            <div class="mb-4">
                <label for="url_start" class="text-gray-600 block mb-2">Đường dẫn</label>
                <input v-model="entry.url_start" type="text" id="url_start" class="w-full border px-4 py-2" required />
            </div>
            <div class="mb-4">
                <label for="avatar" class="text-gray-600 block mb-2">Ảnh</label>
                <input v-model="entry.avatar" type="text" id="avatar" class="w-full border px-4 py-2" required />
            </div>
            <input v-model="entry.typeSource" type="hidden" id="typeSource" />
            <button type="submit" class="w-full bg-blue-500 text-white px-4 py-2 rounded">Thêm</button>
        </form>
    </div>
</template>

  
<script>
import axios from 'axios';

export default {
    name: 'createLabel',
    data() {
        return {
            entry: {
                typeSource: this.$route.query.typeSource || '',
            },
        };
    },
    methods: {
        async createSource() {
            try {
                await axios.post('http://localhost:6006/api/v1/common/sources', this.entry);
                this.$router.push({ path: `/source` });
            } catch (error) {
                console.error('Error adding source:', error);
            }

        },
    },
};
</script>
  
<style></style>
  