<script setup lang="ts">
//base
import { ref, watch } from 'vue'
import { useMessage, NDrawer, NDrawerContent, NInput, NButton, useDialog, NIcon } from 'naive-ui';
const show = defineModel<boolean>('show', { required: true })
const ideaID = defineModel<number>("ideaID", { required: true })
const { placement } = defineProps(["placement"])
//utils
import { Delete as delicon } from '@vicons/carbon'
import { dateToDescription } from '@/utils/time';
import type { CommentItem } from '@/utils/interface'
//store
const Dialog = useDialog()
//api
import { CreateComment, DeleteComment, GetComment } from '@/api/comment';

//ref
const commentList = ref<CommentItem[]>([])
const commentCount = ref<number>(0)
const commentInput = ref<string>("")
const targetCommentID = ref<number>(0)
const dialogOptions = {
    title: '删除',
    content: '是否删除这条评论?',
    positiveText: '确定',
    negativeText: '不确定',
    onPositiveClick: async () => {
        await deleteComment()
    },
    onNegativeClick: () => {
    }
}
//hook
watch(show, async (newValue) => {
    if (newValue) {
        await getComment()
    }
})
//fn
const getComment = async () => {
    console.log(ideaID.value)
    let res = await GetComment({ "IdeaID": ideaID.value })
    if (res.ok) {
        commentCount.value = res.data.CommentCount
        commentList.value = res.data.CommentList
    }
}

const doComment = async () => {
    let content = commentInput.value
    await CreateComment({
        "IdeaID": ideaID.value,
        "Content": content,
    })
    await getComment()
}

const handleDelete = (commentID: number) => {
    targetCommentID.value = commentID
    Dialog.warning(dialogOptions)
}

const deleteComment = async () => {
    await DeleteComment({ "CommentID": targetCommentID.value })
    await getComment()
}
</script>
<template>
    <n-drawer v-model:show="show" :placement="placement" default-width="500" default-height="600" resizable>
        <n-drawer-content :title="'评论💡 ( ' + commentCount + ' ) '">
            <div v-for="comment in commentList" class="comment-card">
                <div class="comment-card-header">
                    <div class="comment-card-author" :style="{ background: comment.UserColor, color: '#fff' }">
                        {{ comment.UserName }}
                    </div>
                    <div style="display: flex;">
                        <div>
                            {{ dateToDescription(comment.CreatedAt) }}
                        </div>
                        <div class="comment-card-del" @click="handleDelete(comment.ID)">
                            <n-icon>
                                <delicon />
                            </n-icon>
                        </div>
                    </div>
                </div>
                <div class="comment-card-content">
                    {{ comment.Content }}
                </div>
            </div>
            <template #footer>
                <div>
                    <n-input v-model:value="commentInput" placeholder="细说细说" type="textarea" clearable
                        maxlength="120"></n-input>
                    <div style="display: flex; margin-top: 12px;">
                        <div style="flex:1"></div>
                        <div style="width: 10px;"></div>
                        <n-button @click="doComment">评论</n-button>
                    </div>
                </div>
            </template>
        </n-drawer-content>
    </n-drawer>
</template>
<style></style>