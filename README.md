# Campus Hub Server

This is the backend service for the campus-hub online. The aim of campus-hub is offering an public lecture resource sharing platform. You can choose to public your lecture resource for multiple formats including `.pdf`, `.avi`, `ppt` etc. The cloud AI Assist would reading those information and help customers learning effectively.

## Develop Requirements







### Lecture Service





> 在教育行业，教学体系的组织通常涉及以下层次的结构，这些结构可以有不同的命名，但一般都包含了类似的概念：
>
> 1. **教学大类（Disciplines）**：这是最高层次的分类，涵盖广泛的学科领域。例如，工程学（Engineering）、物理学（Physics）、计算机学科（Computer Science）等。
> 2. **学科类别（Departments）**：在每个教学大类下，有更具体的学科类别。例如，在计算机学科中，可以有软件工程、数据科学等学科类别。
> 3. **专业（Programs）**：在学科类别下，可以进一步细分为专业。专业通常是一组相关的课程，构建了学生在特定领域的专业知识。
> 4. **课程（Courses or Lectures）**：在专业中，课程是教学的基本单位。每门课程通常涵盖特定的主题，教学内容，并为学生提供相关的学分。有时，“Lecture” 一词用于表示一节课程中的单一教学讲座。
> 5. **章节（Chapters）**：在每门课程内部，内容可能会被组织为不同的章节，每个章节涵盖特定的主题或概念。这样的组织有助于学生更好地理解和组织学科知识。
>
> 这是一个一般性的组织结构，具体的学校或教育机构可能会有一些差异。同时，一些机构可能使用不同的术语，例如，有些地方可能称之为 "Modules" 而不是 "Chapters"。总体而言，这种层次结构有助于组织教育资源、为学生提供清晰的学习路径，并确保教学内容的系统性和逻辑性。

因此我们将这里的层级划分为以下几个层级：

> Disciplines >> Courses >> Chapters

舍弃 Department 和 Programs 的原因在于这两个级别会导致整个层级过于复杂，我将通过 Tag 的方式来实现这两个层级的标注，这样也利于交叉学科融合学科的标明。



![image-20231204114123795](D:\Project\campus-hub\campus-hub-server\assets\image-20231204114123795.png)

![image-20231204114002161](D:\Project\campus-hub\campus-hub-server\assets\image-20231204114002161.png)

```go
type (
	Lecture struct {
		ID   int64
		Name string
		// Publisher The contribution Team or person.
		// TODO create struct for it.
		Publisher string
		License   string
		// LectureVersion, defined the version of lecture
		// example: `23 fall`, `24 spring`, `2023`, `2024`.
		Version string
		TimestampedModel

		// Lecture Description, including four parts.
        // only `Overview` is required, others are optional.
		// LectureDesc struct {
		// 	 Overview        string
		//	 Prerequisites   string
		//	 CourseGoals     string
		//	 FollowupCourses string
		// }

		ChapterList []Chapter
		// TODO Lecture Notice
		// NoticeList  []Notice
	}

	Chapter struct {
		ChapterID          int64
		ChapterName        string
		ResourceCollection []Resource
	}
)

```



#### Changelog

[Vitepress中使用markdown渲染时间轴 - 掘金 (juejin.cn)](https://juejin.cn/post/7237007450178814011)



### Configuration Ref

[fighting-design/README.zh-CN.md at master · FightingDesign/fighting-design (github.com)](https://github.com/FightingDesign/fighting-design/blob/master/README.zh-CN.md)
[快速上手 | Fighting Design (tianyuhao.cn)](https://fighting.tianyuhao.cn/docs/import.html)
[快速开始 | cz-git (qbb.sh)](https://cz-git.qbb.sh/zh/guide/)
[Zhengqbbb/cz-git: cz-git | czg 🛠️ DX first and more engineered, lightweight, customizable, standard output format Commitizen adapter and CLI (github.com)](https://github.com/Zhengqbbb/cz-git)
[LetTTGACO/elog: Markdown 批量导出工具、开放式跨平台博客解决方案，随意组合写作平台(语雀/Notion/FlowUs/飞书)和博客平台(Hexo/Vitepress/Halo/Confluence/WordPress等) (github.com)](https://github.com/LetTTGACO/elog)
[快速开始 | Yuque-VitePress](https://yuque-vitepress.vercel.app/docs/入门指引/快速开始)
[yuque-vitepress/utils/assists.ts at master · LetTTGACO/yuque-vitepress (github.com)](https://github.com/LetTTGACO/yuque-vitepress)





[十分钟教会你如何使用VitePress搭建及部署个人博客站点 - 掘金 (juejin.cn)](https://juejin.cn/post/7128088185917145124)

[Material Symbols and Icons - Google Fonts](https://fonts.google.com/icons?icon.query=email)

PWA [Automatic reload | Guide | Vite PWA (vite-pwa-org.netlify.app)](https://vite-pwa-org.netlify.app/guide/auto-update.html#ssr-ssg)

[markdown-it-wikilinks - npm (npmjs.com)](https://www.npmjs.com/package/markdown-it-wikilinks)
