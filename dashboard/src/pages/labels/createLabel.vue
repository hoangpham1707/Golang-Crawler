<template>
    <div class="max-w-md mx-auto mt-10 p-6 bg-white rounded shadow-md">
        <h1 class="text-2xl font-semibold mb-6">Thêm danh mục</h1>
        <form @submit.prevent="createLabel">
            <div class="mb-4">
                <label for="labelName" class="text-gray-600 block mb-2">Tên nhãn</label>
                <input v-model="entry.name" type="text" id="labelName" class="w-full border px-4 py-2" required />
            </div>
            <div class="mb-4">
                <label for="slug" class="text-gray-600 block mb-2">Slug</label>
                <input v-model="entry.slug" type="text" id="slug" class="w-full border px-4 py-2" required />
            </div>
            <button type="submit" class="w-full bg-blue-500 text-white px-4 py-2 rounded">Lưu</button>
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
                name: '',
                slug: ''
            },
        };
    },
    created() {


    },
    methods: {
        async createLabel() {
            try {
                await axios.post('http://localhost:6006/api/v1/common/labels', this.entry);
                this.$router.push({ path: `/label` });


            } catch (error) {
                console.error('Error adding source:', error);
            }
        },
    },
};
</script>
  
<style></style>
  