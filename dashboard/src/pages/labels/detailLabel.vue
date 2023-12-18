<template>
    <div class="max-w-md mx-auto mt-10 p-6 bg-white rounded shadow-md">
        <h1 class="text-2xl font-semibold mb-6">Sửa danh mục</h1>
        <form @submit.prevent="updateLabel">
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
    name: 'detailLabelView',
    data() {
        return {
            entry: {

            },
        };
    },
    created() {
        this.editingLabel();

    },
    methods: {
        async editingLabel() {
            try {
                const response = await axios.get(`http://localhost:6006/api/v1/common/labels/${this.$route.params.entryId}`);
                console.log({ response });
                this.entry = response.data.data.entry;
            } catch (error) {
                console.error('Error fetching category:', error);
            }
        },
        async updateLabel() {
            try {
                await axios.put(`http://localhost:6006/api/v1/common/labels/${this.$route.params.entryId}`, {
                    name: this.entry.name,
                    slug: this.entry.slug
                });
                this.$router.push('/label');

            } catch (error) {
                console.error('Error updating category:', error);
            }
        },
    },
};
</script>
  
<style></style>
  