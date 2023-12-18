<template>
    <!-- Edit form -->
    <div class="max-w-md mx-auto mt-8 p-6 bg-white rounded shadow-md">
        <h1 class="text-2xl font-semibold mb-6">Sửa Link Next</h1>
        <form @submit.prevent="updateUrlNext">

            <div class="flex items-center mb-4">
                <label for="sourceUrlStart" class="text-gray-600 w-1/3">Url Start</label>
                <input v-model="entry.url_start" type="text" id="sourceUrlStart" class="w-2/3 border px-4 py-2" required />
            </div>

            <div class="flex items-center mb-4">
                <label for="boxElement" class="text-gray-600 w-1/3">Box Element</label>
                <input v-model="entry.box_element" type="text" id="boxElement" class="w-2/3 border px-4 py-2" required />
            </div>

            <div class="flex items-center mb-4">
                <label for="titleElement" class="text-gray-600 w-1/3">TitleElement</label>
                <input v-model="entry.title_element" type="text" id="titleElement" class="w-2/3 border px-4 py-2"
                    required />
            </div>

            <div class="flex items-center mb-4">
                <label for="linkElement" class="text-gray-600 w-1/3">LinkElement</label>
                <input v-model="entry.link_element" type="text" id="sourceName" class="w-2/3 border px-4 py-2" required />
            </div>
            <div class="flex items-center mb-4">
                <label for="sourceName" class="text-gray-600 w-1/3">TimeElement</label>
                <input v-model="entry.time_element" type="text" id="sourceName" class="w-2/3 border px-4 py-2" required />
            </div>
            <div class="flex items-center mb-4">
                <label for="sourceName" class="text-gray-600 w-1/3">DescriptionElement</label>
                <input v-model="entry.description_element" type="text" id="sourceName" class="w-2/3 border px-4 py-2"
                    required />
            </div>
            <div class="flex items-center mb-4">
                <label for="sourceName" class="text-gray-600 w-1/3">CategoryElement</label>
                <input v-model="entry.category_element" type="text" id="sourceName" class="w-2/3 border px-4 py-2"
                    required />
            </div>
            <div class="flex items-center mb-4">
                <label for="sourceName" class="text-gray-600 w-1/3">CheckDesc</label>
                <input v-model="entry.content_element" type="text" id="sourceName" class="w-2/3 border px-4 py-2"
                    required />
            </div>
            <div class="flex items-center mb-4">
                <label for="sourceName" class="text-gray-600 w-1/3">ContentElement</label>
                <input v-model="entry.check_time" type="text" id="sourceName" class="w-2/3 border px-4 py-2" required />
            </div>
            <div class="flex items-center mb-4">
                <label for="sourceName" class="text-gray-600 w-1/3">CheckDes</label>
                <input v-model="entry.check_desc" type="text" id="sourceName" class="w-2/3 border px-4 py-2" required />
            </div>
            <div class="flex items-center mb-4">
                <label for="sourceName" class="text-gray-600 w-1/3">CheckCategory</label>
                <input v-model="entry.check_category" type="text" id="sourceName" class="w-2/3 border px-4 py-2" required />
            </div>

            <button type="submit" class="bg-blue-500 text-white px-2 py-1 rounded">Lưu</button>
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
        this.editingUrlNext();
    },
    methods: {
        async editingUrlNext() {
            try {
                const response = await axios.get(`http://localhost:6006/api/v1/common/linkNexts/${this.$route.params.entryId}`);
                console.log({ response });
                this.entry = response.data.data.entry;
            } catch (error) {
                console.error('Error fetching category:', error);
            }
        },
        async updateUrlNext() {
            try {
                await axios.put(`http://localhost:6006/api/v1/common/linkNexts/${this.$route.params.entryId}`, {
                    url_start: this.entry.url_start,
                    box_element: this.entry.box_element,
                    title_element: this.entry.title_element,
                    link_element: this.entry.link_element,
                    time_element: this.entry.time_element,
                    description_element: this.entry.description_element,
                    category_element: this.entry.category_element,
                    content_element: this.entry.content_element,
                    check_time: this.entry.check_time,
                    check_desc: this.entry.check_desc,
                    check_category: this.entry.check_category,
                });
                this.$router.push('/urlNext');

            } catch (error) {
                console.error('Error updating category:', error);
            }
        },
    },
};
</script>
  
<style></style>
  