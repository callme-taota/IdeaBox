<script setup lang="ts">
//base
import { ref } from 'vue'
import { useMessage, NDrawer, NDrawerContent, NInput, NButton } from 'naive-ui';
const emit = defineEmits(['done'])
const show = defineModel<boolean>('show', { required: true })
const { placement } = defineProps(["placement"])
//utils
import type { Idea } from '@/utils/interface'
//store
const Message = useMessage()
//api
import { CreateIdea } from '@/api/idea';

//ref
const ideaItem = ref<Idea>({
    Name: '',
    Description: '',
    Author: 1
})
//hook

//fn
const submitIdea = async () => {
    let obj = ideaItem.value
    if (obj.Name == "" || obj.Description == "") {
        Message.warning("看看是不是少写了些啥")
    }
    await CreateIdea(obj)
    show.value = false
    ideaItem.value = {
        Name: "",
        Description: "",
        Author: ideaItem.value.Author,
    }
    emit('done')
}
</script>
<template>
    <n-drawer v-model:show="show" :placement="placement" default-width="500" default-height="350">
        <n-drawer-content title="我有一个点子💡">
            <div>
                点子
                <n-input v-model:value="ideaItem.Name" placeholder="点子" maxlength="30" show-count clearable></n-input>
            </div>
            <div>
                细说细说
                <n-input v-model:value="ideaItem.Description" placeholder="细说细说" type="textarea" clearable></n-input>
            </div>
            <n-button @click="submitIdea">想完了</n-button>
        </n-drawer-content>
    </n-drawer>
</template>
<style></style>