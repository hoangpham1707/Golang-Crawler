<template>
    <!-- Edit form -->
    <div class="max-w-md mx-auto mt-8 p-6 bg-white rounded shadow-md">
        <h1 class="text-2xl font-semibold mb-6">Sửa nguồn đăng tải</h1>
        <form @submit.prevent="updateSource">
            <div class="mb-4">
                <label for="sourceName" class="text-gray-600 block mb-2">Đường dẫn</label>
                <input v-model="entry.url_start" type="text" id="sourceName" class="w-full border px-4 py-2" required />
            </div>

            <div class="mb-4">
                <label for="avatar" class="text-gray-600 block mb-2">Ảnh</label>
                <input v-model="entry.avatar" type="text" id="avatar" class="w-full border px-4 py-2" required />
            </div>

            <button type="submit" class="w-full bg-blue-500 text-white px-4 py-2 rounded">Lưu</button>
        </form>
    </div>
</template>

  
<script>
import axios from 'axios';

export default {
    name: 'DetailSource',
    data() {
        return {
            entry: {

            },
        };
    },
    created() {
        this.editingSource();
    },
    methods: {
        async editingSource() {
            try {
                const response = await axios.get(`http://localhost:6006/api/v1/common/sources/${this.$route.params.entryId}`);
                console.log({ response });
                this.entry = response.data.data.entry;
            } catch (error) {
                console.error('Error fetching category:', error);
            }
        },
        async updateSource() {
            try {
                await axios.put(`http://localhost:6006/api/v1/common/sources/${this.$route.params.entryId}`, {
                    url_start: this.entry.url_start,
                    avatar: this.entry.avatar
                });

                const typeSource = this.$route.query.typeSource || '';
                this.$router.push({
                    path: `/typeSource`,
                    query: { typeSource }
                });

            } catch (error) {
                console.error('Error updating category:', error);
            }
        },
    },
};
</script>
  
<style></style>
  