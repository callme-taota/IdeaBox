<script setup lang="ts">
import { ref, onMounted, onUnmounted, Transition } from 'vue';
import { NIcon } from 'naive-ui'
import { Sunny, Moon, SyncOutline } from '@vicons/ionicons5'
import { TwoFactorAuthentication as authicon } from '@vicons/carbon'
import { useThemeStore } from '@/stores/theme'
import { useAuthStore } from '@/stores/auth';
import { storeToRefs } from 'pinia';

const isScrolled = ref(false)
const themeChanger = ref(false)
const themeStore = useThemeStore()
const authStore = useAuthStore()
const { showAuthModal } = storeToRefs(authStore)

const handleScroll = () => {
    isScrolled.value = window.scrollY > 0;
};

onMounted(() => {
    themeStore.SetThemeAuto()
    window.addEventListener('scroll', handleScroll);
});

onUnmounted(() => {
    window.removeEventListener('scroll', handleScroll);
});

const setTheme = (themeID: number) => {
    if (themeID == 1) {
        themeStore.SetThemeAuto()
    } else if (themeID == 2) {
        themeStore.SetThemeLight()
    } else {
        themeStore.SetThemeDark()
    }
    hideThemeChanger()
}

const showThemeChanger = () => {
    themeChanger.value = !themeChanger.value
}

const hideThemeChanger = () => {
    themeChanger.value = false
}

const showAuth = () => {
    showAuthModal.value = true
}
</script>
<template>
    <Transition>
        <div class="layout-header-cont" :class="{ 'layout-header-cont-after-scroll': isScrolled }">
            <div class="layout-header-left">
                <div class="layout-header-painter">点子箱💡</div>
            </div>
            <div class="layout-header-center">

            </div>
            <div class="layout-header-right">
                <div class="layout-header-right-icon-cont" @click="showThemeChanger">
                    <n-icon size="20">
                        <sunny />
                    </n-icon>
                </div>
                <div class="layout-header-right-icon-cont" @click="showAuth">
                    <n-icon size="20">
                        <authicon />
                    </n-icon>
                </div>
            </div>
            <Transition name="fade-slide">
                <div class="layout-header-theme-changer-cont" v-show="themeChanger" @mouseleave="hideThemeChanger">
                    <div class="layout-header-theme-changer-item" @click="setTheme(2)">
                        <n-icon size="20">
                            <sunny />
                        </n-icon>
                        &nbsp;&nbsp;
                        浅色
                    </div>
                    <div class="layout-header-theme-changer-item" @click="setTheme(3)">
                        <n-icon size="20">
                            <moon />
                        </n-icon>
                        &nbsp;&nbsp;
                        深色
                    </div>
                    <div class="layout-header-theme-changer-item" @click="setTheme(1)">
                        <n-icon size="20">
                            <sync-outline />
                        </n-icon>
                        &nbsp;&nbsp;
                        跟随系统
                    </div>
                </div>
            </Transition>
        </div>
    </Transition>
</template>
<style>
.layout-header-cont {
    width: calc(100% - 100px);
    height: 60px;
    display: flex;
    position: fixed;
    padding: 0 50px;
    justify-content: space-between;
    transition: 0.5s;
    z-index: 2;
}

.layout-header-cont-after-scroll {
    background-color: var(--header-background);
    backdrop-filter: blur(12px);
}

.layout-logo {
    height: 40px;
    border-radius: 14px;
    transition: 0.3s;
}

.layout-logo:hover {
    background-color: var(--base-hover-background);
}

.layout-header-left {
    display: flex;
    padding: 0 10px;
    align-items: center;
}

.layout-header-painter {
    margin-left: 10px;
    line-height: 40px;
    height: 40px;
    border-radius: 14px;
    transition: 0.3s;
    padding: 0 10px;
    user-select: none;
    font-weight: bolder;
    font-size: large
}

.layout-header-painter:hover {
    background-color: var(--base-hover-background);
    color: var(--btn-hover-color);
}

.layout-header-right {
    display: flex;
    align-items: center;
    flex-direction: row-reverse;
}

.layout-header-right-icon-cont {
    transition: 0.3s;
    margin: 0 4px;
    border-radius: 14px;
    height: 40px;
    width: 40px;
    display: flex;
    align-items: center;
    cursor: pointer;
    justify-content: center;
}

.layout-header-right-icon-cont:hover {
    background-color: var(--base-hover-background);
    color: var(--btn-hover-color);
}

.layout-header-right-text-cont {
    margin: 0 4px;
    transition: 0.3s;
    border-radius: 14px;
    height: 40px;
    padding: 0 12px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: bold;
    user-select: none;
    cursor: pointer;
}

.layout-header-right-text-cont:hover {
    background-color: var(--base-hover-background);
    color: var(--btn-hover-color);
}

.layout-header-theme-changer-cont {
    position: fixed;
    background-color: var(--header-background);
    backdrop-filter: blur(12px);
    border-radius: 16px;
    border: 1px solid var(--border-color);
    width: 100px;
    height: 110px;
    top: 60px;
    right: 22px;
    padding: 2px 6px;
    transition: 0.5s;
}

.layout-header-theme-changer-item {
    height: 30px;
    line-height: 30px;
    display: flex;
    margin: 5px 0;
    justify-content: start;
    align-items: center;
    user-select: none;
    padding-left: 6px;
    border-radius: 8px;
    cursor: pointer;
}

.layout-header-theme-changer-item:hover {
    background-color: var(--btn-hover-grey);
}

.latout-header-right-icon-afterhide {
    display: none;
}

@media (max-width: 1200px) {
    .layout-header-right-text-cont {
        display: none;
    }

    .latout-header-right-icon-afterhide {
        display: flex;
    }

    .layout-header-cont {
        padding: 0 20px;
        width: calc(100% - 40px);
    }
}
</style>