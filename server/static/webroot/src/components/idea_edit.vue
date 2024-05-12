<script setup lang="ts">
//base
import { useMessage, NDrawer, NDrawerContent, NInput, NButton, NSelect } from 'naive-ui';
const emit = defineEmits(['done'])
const show = defineModel<boolean>('show', { required: true })
const editItem = defineModel<IdeaWithID>('editItem', { required: true })
const { placement, userList } = defineProps(["placement","userList"])
//utils
import type { IdeaWithID } from '@/utils/interface'
//store
const Message = useMessage()
//api
import { UpdateIdea, DeleteIdea } from '@/api/idea';

//ref
//hook

//fn
const updateIdea = async () => {
    let obj = editItem.value
    if (obj.Name == "" || obj.Author > 2 || obj.Author < 1 || obj.Description == "") {
        Message.warning("看看是不是少写了些啥")
    }
    await UpdateIdea(obj)
    show.value = false
    editItem.value = {
        ID: 0,
        Name: "",
        Description: "",
        Author: editItem.value.Author,
    }
    emit('done')
}

const delIdea = async () => {
    let id = editItem.value.ID
    await DeleteIdea({
        "ID": id,
    })
    show.value = false
    editItem.value = {
        ID: 0,
        Name: "",
        Description: "",
        Author: editItem.value.Author,
    }
    emit('done')
}
</script>
<template>
    <n-drawer v-model:show="show" :placement="placement" default-width="500" default-height="370">
        <n-drawer-content title="点子得改改💡">
            <div>
                点子
                <n-input v-model:value="editItem.Name" placeholder="点子" maxlength="30" show-count clearable></n-input>
            </div>
            <div>
                谁
                <n-select v-model:value="editItem.Author" placeholder="谁呀" :options="userList" label-field="Name"
                    value-field="ID" filterable></n-select>
            </div>
            <div>
                细说细说
                <n-input v-model:value="editItem.Description" placeholder="细说细说" type="textarea" clearable></n-input>
            </div>
            <br>
            <div style="display: flex; justify-content: space-between;">
                <n-button @click="updateIdea">想完了</n-button>
                <n-button @click="delIdea" type="error">删了它</n-button>
            </div>
        </n-drawer-content>
    </n-drawer>
</template>
<style></style>