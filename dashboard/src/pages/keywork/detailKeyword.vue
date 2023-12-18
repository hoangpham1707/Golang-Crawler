<template>
    <div>
        <h1>Sửa từ khóa</h1>
        <form @submit.prevent="updateKeyword">
            <div class="mb-4">
                <label for="keywordName" class="text-gray-600 w-1/3">Từ khóa</label>
                <input v-model="entry.name" type="text" id="keywordName" class="w-2/3 border px-4 py-2" required />
            </div>
            <button type="submit" class="w-full bg-blue-500 text-white px-4 py-2 rounded">Lưu</button>
        </form>
    </div>
</template>
  
<script>
import axios from 'axios';

export default {
    name: 'detailKeywordView',
    data() {
        return {
            entry: {

            },
        };
    },
    created() {
        this.editingKeyword();
    },
    methods: {
        async editingKeyword() {
            try {
                const response = await axios.get(`http://localhost:6006/api/v1/common/keywords/${this.$route.params.entryId}`);
                console.log({ response });
                this.entry = response.data.data.entry;
            } catch (error) {
                console.error('Error fetching category:', error);
            }
        },
        async updateKeyword() {
            try {
                await axios.put(`http://localhost:6006/api/v1/common/keywords/${this.$route.params.entryId}`, {
                    name: this.entry.name,
                });
                this.$router.push('/keyword');

            } catch (error) {
                console.error('Error updating category:', error);
            }
        },
    },
};
</script>
  
<style></style>
  