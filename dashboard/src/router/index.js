import { createRouter, createWebHistory } from "vue-router";
import dashboard from '../pages/master/dashboard.vue'
import home from '../pages/home.vue';
import source from '../pages/sources/source.vue'
import urlNext from '../pages/link/urlNext.vue'
import category from '../pages/category/category.vue'
import keyword from '../pages/keywork/keyword.vue'
import addLabel from '../pages/labels/addLabel.vue'
import label from '../pages/labels/label.vue'
import sentences from '../pages/labels/sentences.vue'
import detailCategory from '../pages/category/detailCategory.vue'
import detailSource from '../pages/sources/detailSource.vue'
import detailKeyword from '../pages/keywork/detailKeyword.vue'
import detailLinkNext from '../pages/link/detailUrlNext.vue'
import detailLabel from '../pages/labels/detailLabel.vue'
import createLabel from '../pages/labels/createLabel.vue'
import createSource from '../pages/sources/createSource.vue'
import createLinkNext from '../pages/link/createLinkNext.vue'
import addLabelWeb from '../pages/labels/addLabelWeb.vue'
import face from '../pages/facebook/pages.vue'
import typeSource from '../pages/sources/newSource.vue'


const routes = [
    {
        name: 'DashBoard',
        path: '/',
        component: dashboard,
        children: [
            {
                name: 'Home',
                path: '/',
                component: home,
            },
            {
                name: 'ListSource',
                path: '/source',
                component: source,
            },
            {
                name: 'UrlNext',
                path: '/urlNext',
                component: urlNext,
            },
            {
                name: 'createLinkNext',
                path: '/createLinkNext',
                component: createLinkNext,
            },
            {
                name: 'DetailUrlNext',
                path: '/detailLinkNext/:entryId',
                component: detailLinkNext,
            },
            {
                name: 'categoryView',
                path: '/category',
                component: category,
            },
            {
                name: 'CategoryView',
                path: '/detailCategory/:entryId',
                component: detailCategory,
            },
            {
                name: 'SourceView',
                path: '/detailSource/:entryId',
                component: detailSource,
            },
            {
                name: 'CreateSource',
                path: '/createSource',
                component: createSource,
            },
            {
                name: 'TypeSource',
                path: '/typeSource',
                component: typeSource,
            },
            {
                name: 'keywordView',
                path: '/keyword',
                component: keyword,
            },
            {
                name: 'detailKeywordView',
                path: '/detailKeyword/:entryId',
                component: detailKeyword,
            },
            {
                name: 'AddLabelView',
                path: '/addLabels',
                component: addLabel,
            },
            {
                name: 'facePages',
                path: '/face',
                component: face,
            },
            {
                name: 'AddLabelWebView',
                path: '/web',
                component: addLabelWeb,
            },
            {
                name: 'LabelView',
                path: '/label',
                component: label,
            },
            {
                name: 'detailLabel',
                path: '/detailLabel/:entryId',
                component: detailLabel,
            },
            {
                name: 'createLabel',
                path: '/createLabel',
                component: createLabel,
            },
            {
                name: 'SentencesView',
                path: '/sentences/:entryId',
                component: sentences,
            }
        ]
    },
];

const router = Router();
export default router;

function Router() {
    const router = createRouter({
        history: createWebHistory(),
        routes,
    })
    return router;
}
