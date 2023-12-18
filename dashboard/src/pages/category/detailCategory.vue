<template>
    <!-- Edit form -->
    <div>
        <h1>Sửa danh mục</h1>
        <form @submit.prevent="updateCategory">
            <div class="mb-4">
                <label for="categoryName" class="text-gray-600 w-1/3">Tên danh mục</label>
                <input v-model="entry.name" type="text" id="categoryName" class="w-2/3 border px-4 py-2" required />
            </div>
            <button type="submit" class="w-full bg-blue-500 text-white px-4 py-2 rounded">Lưu</button>
        </form>
    </div>
</template>
  
<script>
import axios from 'axios';

export default {
    name: 'DetailCategory',
    data() {
        return {
            entry: {

            },
        };
    },
    created() {
        this.editingCategory();
    },
    methods: {
        async editingCategory() {
            try {
                const response = await axios.get(`http://localhost:6006/api/v1/common/categories/${this.$route.params.entryId}`);
                console.log({ response });
                this.entry = response.data.data.entry;
            } catch (error) {
                console.error('Error fetching category:', error);
            }
        },
        async updateCategory() {
            try {
                await axios.put(`http://localhost:6006/api/v1/common/categories/${this.$route.params.entryId}`, {
                    name: this.entry.name,
                });
                this.$router.push('/category');

            } catch (error) {
                console.error('Error updating category:', error);
            }
        },
    },
};
</script>
  
<style></style>
  