<script setup lang="ts">
import { dateToString } from '@/utils/time';
import { CalendarEdit24Regular, Comment20Regular } from '@vicons/fluent'
import { NIcon } from 'naive-ui';
const { userName, ideaName, ideaDescription, createAt, color } = defineProps<{
    userName: string,
    ideaName: string,
    ideaDescription: string,
    createAt: string,
    color: string
}>()
const emit = defineEmits(['edit','comment']);
const handleEditClick = () => {
    emit('edit');
};
const handleCommentClick = () => {
    emit('comment');
};
</script>
<template>
    <div class="idea-card-cont">
        <div class="idea-card-head">
            <div class="idea-card-title">
                {{ ideaName }}
            </div>
            <div style="display: flex; align-items: center;">
                <div class="idea-card-author" :style="{ background: color, color: '#fff' }">
                    {{ userName }}
                </div>
                <div class="idea-edit-btn" @click="handleCommentClick">
                    <n-icon>
                        <Comment20Regular />
                    </n-icon>
                </div>
                <div class="idea-edit-btn" @click="handleEditClick">
                    <n-icon>
                        <CalendarEdit24Regular />
                    </n-icon>
                </div>
            </div>
        </div>
        <div class="idea-card-desc">
            {{ ideaDescription }}
        </div>
        <div class="idea-card-time">
            {{ dateToString(createAt) }}
        </div>
    </div>
</template>
<style>
.idea-card-cont {
    border-radius: 14px;
    padding: 16px;
    background: var(--card--background);
    margin-bottom: 10px;
    transition: 0.3s;
    break-inside: avoid;
}

.idea-card-cont:hover {
    box-shadow: 0 0 4px var(--highlight-color);
}

.idea-card-head {
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.idea-card-title {
    font-weight: bold;
    font-size: larger;
}

.idea-card-author {
    margin-left: 10px;
    padding: 1px 6px;
    font-weight: 300;
    border-radius: 6px;
    margin-right: 6px;
}

.idea-edit-btn {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 24px;
    height: 24px;
    border-radius: 50%;
    background: var(--add-btn-background);
    cursor: pointer;
    transition: 0.3s;
    margin: 0 4px;
}

.idea-edit-btn:hover {
    background: var(--second-highlight-color);
    color: var(--color-rev);
}

.idea-card-desc {
    margin-top: 6px;
    padding: 6px;
    border-top: var(--content-border);
    overflow-wrap: break-word;
}

.idea-card-time {
    margin-top: 12px;
    font-weight: 300;
    font-size: small;
}
</style>