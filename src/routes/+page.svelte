<script>
  const contact = [
    { icon: 'bx bx-envelope', label: 'melkir13@gmail.com', href: 'mailto:melkir13@gmail.com' },
    { icon: 'bx bx-globe', label: 'thibaultvieux.com', href: 'https://thibaultvieux.com/' },
    { icon: 'bx bxl-github', label: 'github.com/melkir', href: 'https://github.com/melkir' },
    {
      icon: 'bx bxl-linkedin',
      label: 'linkedin.com/in/thibault-vieux',
      href: 'https://www.linkedin.com/in/thibault-vieux'
    },
    { icon: 'bx bx-map', label: 'Japan', href: null }
  ];

  const skills = {
    languages: [
      'JavaScript (ES2015+)',
      'TypeScript',
      'HTML',
      'CSS',
      'Rust',
      'SQL',
      'GraphQL',
      'Python'
    ],
    frameworks: [
      'React',
      'Next.js',
      'Node.js',
      'Angular',
      'Tailwind CSS',
      'Framer Motion',
      'Postgres',
      'MongoDB',
      'Convex'
    ],
    tools: [
      'Git',
      'GitHub',
      'Linear',
      'ElasticSearch',
      'Docker',
      'Kubernetes',
      'Terraform',
      'MongoDB',
      'Postgres',
      'Vercel',
      'AWS',
      'Figma'
    ]
  };

  const projects = [
    {
      name: 'Remote GPIO',
      description:
        'Rust-based GPIO control system for remote hardware management with web interface',
      url: 'https://github.com/melkir/remote-gpio'
    },
    {
      name: 'Pentest Workshop',
      description:
        'Educational penetration testing workshop materials and comprehensive security tools',
      url: 'https://github.com/melkir/hacking-workshop'
    },
    {
      name: 'SMS Push',
      description: 'Android SMS forwarding system for remote message management and notifications',
      url: 'https://github.com/melkir/android-sms-push'
    }
  ];
  const experience = [
    {
      title: 'Technical Consultant',
      company: 'Freelance',
      location: 'Remote',
      period: 'Sept 2024 — Present',
      details: [
        'Operate as an independent consultant providing expert solutions in technical architecture, system design, and full-stack development.',
        'Developed an automated compliance system at EarlyWaters, utilizing confidential computing for tokenized RWA reporting.',
        'Built a proof-of-concept map-based AI agent leveraging RAG/MCP framework for traveler recommendations.'
      ]
    },
    {
      title: 'Senior Backend Software Engineer',
      company: 'Cryptio',
      companyUrl: 'https://cryptio.co/',
      location: 'Paris',
      period: 'April — July 2023',
      details: [
        'Led backend team in transforming digital asset data into auditable financial records using NestJS, React and TypeScript.',
        'Optimized PostgreSQL database performance for high-volume crypto accounting, reducing query latency.'
      ]
    },
    {
      company: 'PayFit',
      companyUrl: 'https://payfit.com/',
      location: 'Paris',
      positions: [
        {
          title: 'DevOps Engineer',
          period: 'Jan 2021 — Mar 2023'
        },
        {
          title: 'Fullstack Software Engineer',
          period: 'Apr 2018 — Dec 2020'
        }
      ],
      details: [
        'Managed and scaled Kubernetes clusters, enhancing platform reliability and automating CI/CD pipelines.',
        'Implemented Infrastructure as Code (IaC) using Terraform, automating environment provisioning and reducing manual work by 80%.',
        'Improved system observability and reduced incident response time using Datadog monitoring.',
        'Developed internal tools that automated critical government compliance tasks, saving 20+ hours per week.',
        'Architected high-performance GraphQL dashboards with fuzzy search for managing massive JSON datasets.',
        'Engineered real-time data synchronization pipeline between MongoDB and Elasticsearch using event-driven CDC architecture.'
      ]
    },
    {
      title: 'Fullstack Developer',
      company: 'Heek',
      location: 'Paris',
      period: 'Apr — Oct 2017',
      details: [
        'Built a conversational AI chatbot that generated custom websites based on user input.',
        'Collaborated within an agile team to ship features and iterate on the product.'
      ]
    },
    {
      title: 'Mobile Application Developer',
      company: 'Export Japan',
      companyUrl: 'https://www.export-japan.co.jp/',
      location: 'Osaka',
      period: 'May — Sep 2016',
      details: [
        'Implemented a BaaS server to enable secure multi-account management for a hybrid mobile application.',
        'Built a cross-platform mobile app utilizing smartphone capabilities including GPS, NFC, and haptic feedback.'
      ]
    }
  ];

  const education = [
    {
      degree: 'Master in Computer Science',
      school: 'CY Tech',
      period: '2015 — 2017',
      concentration: 'Smart Systems',
      details:
        'Valedictorian of engineering school. Returned to teach penetration testing and network security after graduating.'
    },
    {
      degree: 'Bachelor in Computer Science',
      school: 'Université Paris Cité',
      period: '2013 — 2015',
      concentration: 'Artificial Intelligence',
      details:
        'Gained proficiency in Java and Python with comprehensive coursework in algorithms, mathematics, and machine learning.'
    }
  ];

  const languages = [
    { language: 'French', level: 'Native' },
    { language: 'English', level: 'Fluent' },
    { language: 'Japanese', level: 'Conversational' }
  ];

</script>

<!-- Print optimized resume with proper two-column layout -->
<div class="min-h-screen bg-gray-50 py-8 px-4 print:bg-white print:p-0 print:m-0">
  <div
    class="max-w-5xl mx-auto bg-white shadow-lg rounded-lg print:shadow-none print:rounded-none print:max-w-none"
  >
    <div class="p-12 print:p-0 print:max-h-none">
      <!-- Header -->
      <header class="mb-8 print:mb-6">
        <div class="grid grid-cols-3 gap-6 print:gap-4 items-center">
          <div class="col-span-2 flex items-center gap-6 print:gap-4">
            <div class="flex-shrink-0">
              <img
                src="/avatar.jpg"
                alt="Thibault Vieux"
                class="w-32 h-32 rounded-lg object-cover shadow-md print:w-24 print:h-24 print:shadow-sm"
              />
            </div>
            <div class="flex-1">
              <h1 class="text-6xl font-bold text-sky-600 mb-2 print:text-4xl print:mb-1">
                Thibault Vieux
              </h1>
              <p class="text-base text-gray-600 print:text-sm">
                Experienced software engineer with a focus on building scalable systems and
                technical architecture
              </p>
            </div>
          </div>
          <div class="col-span-1">
            <div class="space-y-2 print:space-y-1.5 text-xs text-gray-600 print:text-[10px]">
              {#each contact as item}
                {#if item.href}
                  <a href={item.href} class="flex items-center hover:text-sky-600" target="_blank">
                    <i class="{item.icon} text-sm mr-2 w-4 flex justify-center print:text-xs"></i>
                    <span class="break-all">{item.label}</span>
                  </a>
                {:else}
                  <div class="flex items-center">
                    <i class="{item.icon} text-sm mr-2 w-4 flex justify-center print:text-xs"></i>
                    <span>{item.label}</span>
                  </div>
                {/if}
              {/each}
            </div>
          </div>
        </div>
      </header>

      <!-- Two-column layout -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8 print:grid-cols-3 print:gap-6">
        <!-- Main Content (Left - 2/3) -->
        <main class="lg:col-span-2 print:col-span-2 space-y-8 print:space-y-6">
          <!-- Experience -->
          <section class="print:break-inside-avoid">
            <h2
              class="text-sm font-semibold text-sky-600 uppercase tracking-wider mb-4 print:mb-3 print:text-xs"
            >
              Revelent Experience
            </h2>
            <div class="space-y-6 print:space-y-4">
              {#each experience as exp}
                <div class="print:break-inside-avoid">
                  <div class="mb-3 print:mb-2">
                    {#if exp.positions}
                      <!-- Multiple positions at same company -->
                      <h3
                        class="text-sm font-semibold text-gray-900 mb-1 flex flex-wrap items-center gap-1.5 print:text-xs"
                      >
                        <span class="font-normal">{exp.positions[0].title}</span>
                        <span class="text-gray-400 font-normal">·</span>
                        {#if exp.companyUrl}
                          <a href={exp.companyUrl} target="_blank" class="font-medium text-gray-900 hover:text-sky-600 no-underline">{exp.company}</a>
                        {:else}
                          <span class="font-medium">{exp.company}</span>
                        {/if}
                        <span class="text-gray-400 font-normal">·</span>
                        <span class="font-normal">{exp.location}</span>
                        <span class="text-gray-400 font-normal">·</span>
                        <span class="text-slate-400 font-normal">{exp.positions[0].period}</span>
                      </h3>
                      {#each exp.positions.slice(1) as position}
                        <div class="text-sm text-gray-700 print:text-xs">
                          <span class="font-normal">{position.title}</span>
                          <span class="text-slate-400 font-normal ml-2">{position.period}</span>
                        </div>
                      {/each}
                    {:else}
                      <!-- Single position -->
                      <h3
                        class="text-sm font-semibold text-gray-900 mb-1 flex flex-wrap items-center gap-1.5 print:text-xs"
                      >
                        <span class="font-normal">{exp.title}</span>
                        <span class="text-gray-400 font-normal">·</span>
                        {#if exp.companyUrl}
                          <a href={exp.companyUrl} target="_blank" class="font-medium text-gray-900 hover:text-sky-600 no-underline">{exp.company}</a>
                        {:else}
                          <span class="font-medium">{exp.company}</span>
                        {/if}
                        <span class="text-gray-400 font-normal">·</span>
                        <span class="font-normal">{exp.location}</span>
                        <span class="text-gray-400 font-normal">·</span>
                        <span class="text-slate-400 font-normal">{exp.period}</span>
                      </h3>
                    {/if}
                  </div>
                  <ul class="space-y-1 text-sm text-gray-600 print:text-xs print:space-y-0.5">
                    {#each exp.details as detail}
                      <li class="flex items-start">
                        <span class="text-sky-600 mr-2 flex-shrink-0 font-bold">•</span>
                        <span class="leading-relaxed print:leading-normal">{detail}</span>
                      </li>
                    {/each}
                  </ul>
                </div>
              {/each}
            </div>
          </section>
        </main>

        <!-- Sidebar (Right - 1/3) -->
        <aside class="print:col-span-1 space-y-6 print:space-y-4">
          <!-- Skills -->
          <section class="print:break-inside-avoid">
            <h2
              class="text-sm font-semibold text-sky-600 uppercase tracking-wider mb-3 print:mb-2 print:text-xs"
            >
              Skills
            </h2>
            <div class="space-y-3 print:space-y-2">
              <div>
                <h3 class="text-xs font-semibold text-gray-900 mb-1 print:text-[10px]">
                  Programming Languages
                </h3>
                <p
                  class="text-sm text-gray-600 leading-relaxed print:text-xs print:leading-tight"
                >
                  {skills.languages.join(', ')}
                </p>
              </div>
              <div>
                <h3 class="text-xs font-semibold text-gray-900 mb-1 print:text-[10px]">
                  Libraries & Frameworks
                </h3>
                <p
                  class="text-sm text-gray-600 leading-relaxed print:text-xs print:leading-tight"
                >
                  {skills.frameworks.join(', ')}
                </p>
              </div>
              <div>
                <h3 class="text-xs font-semibold text-gray-900 mb-1 print:text-[10px]">
                  Tools & Platforms
                </h3>
                <p
                  class="text-sm text-gray-600 leading-relaxed print:text-xs print:leading-tight"
                >
                  {skills.tools.join(', ')}
                </p>
              </div>
            </div>
          </section>

          <!-- Languages -->
          <section class="print:break-inside-avoid">
            <h2
              class="text-sm font-semibold text-sky-600 uppercase tracking-wider mb-3 print:mb-2 print:text-xs"
            >
              Languages
            </h2>
            <p class="text-sm text-gray-600 leading-relaxed print:text-xs print:leading-tight">
              {#each languages as lang, i}
                <span>{lang.language}</span>
                {lang.level}{#if i < languages.length - 1},&nbsp;{/if}
              {/each}
            </p>
          </section>

          <!-- Projects -->
          <section class="print:break-inside-avoid">
            <h2
              class="text-sm font-semibold text-sky-600 uppercase tracking-wider mb-3 print:mb-2 print:text-xs"
            >
              Selected Projects
            </h2>
            <div class="space-y-3 print:space-y-2">
              {#each projects as project}
                <div class="print:break-inside-avoid">
                  <h3 class="text-xs font-semibold mb-1 print:text-[10px]">
                    <a href={project.url} target="_blank" class="text-medium no-underline">
                      {project.name}
                    </a>
                  </h3>
                  <p
                    class="text-sm text-gray-600 leading-relaxed print:text-xs print:leading-tight"
                  >
                    {project.description}
                  </p>
                </div>
              {/each}
            </div>
          </section>

          <!-- Education -->
          <section class="print:break-inside-avoid">
            <h2
              class="text-sm font-semibold text-sky-600 uppercase tracking-wider mb-3 print:mb-2 print:text-xs"
            >
              Education
            </h2>
            <div class="space-y-3 print:space-y-2">
              {#each education as edu}
                <div class="print:break-inside-avoid">
                  <h3 class="text-xs font-semibold text-gray-800 mb-1 print:text-[10px]">
                    {edu.degree}
                  </h3>
                  <div class="text-sm text-gray-600 mb-1.5 print:text-xs print:mb-1">
                    <div class="font-medium text-gray-900">{edu.school}</div>
                    <div class="text-gray-600">{edu.concentration} · {edu.period}</div>
                  </div>
                  <p
                    class="text-sm text-gray-600 leading-relaxed print:text-xs print:leading-tight"
                  >
                    {edu.details}
                  </p>
                </div>
              {/each}
            </div>
          </section>

        </aside>
      </div>
    </div>
  </div>
</div>
