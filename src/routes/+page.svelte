<script>
  import CalendarIcon from '../components/icons/Calendar.svelte';
  import CryptioLogo from '../components/icons/Cryptio.svelte';
  import EducationIcon from '../components/icons/Education.svelte';
  import ExportJapanLogo from '../components/icons/ExportJapan.svelte';
  import FreelancerLogo from '../components/icons/Freelancer.svelte';
  import HeekLogo from '../components/icons/Heek.svelte';
  import JobIcon from '../components/icons/Job.svelte';
  import MapIcon from '../components/icons/Map.svelte';
  import PayfitLogo from '../components/icons/Payfit.svelte';

  const socials = [
    {
      href: 'mailto:melkir13@gmail.com',
      icon: 'bx bx-envelope'
    },
    {
      href: 'https://github.com/melkir',
      icon: 'bx bxl-github'
    },
    {
      href: 'https://line.me/ti/p/VEfTVnxksB',
      icon: 'bx bx-message'
    },
    {
      href: 'https://www.linkedin.com/in/thibault-vieux',
      icon: 'bx bxl-linkedin'
    }
  ];

  const skills = [
    'TypeScript',
    'Rust',
    'React',
    'Angular',
    'MongoDB',
    'SQL',
    'Docker',
    'GraphQL',
    'Kubernetes',
    'Git',
    'Motion',
    'Next.js',
    'Tanstack',
    'Convex'
  ];

  const projects = [
    { href: 'https://github.com/melkir/next-remote-gpio', name: 'Remote GPIO' },
    { href: 'https://github.com/melkir/hacking-workshop', name: 'Pentest Workshop' },
    { href: 'https://github.com/melkir/android-sms-push', name: 'SMS Push' },
    { href: 'https://github.com/melkir/API-Libraries-Demo', name: 'API Libraries Demo' }
  ];

  const experience = [
    {
      title: 'Technical Consultant',
      company: 'Freelance',
      location: 'Remote',
      date: 'Sept 2024 – Present',
      type: 'Freelance',
      duration: '1 year',
      description:
        'Founded a technical consulting startup providing architectural guidance and system design services to early-stage companies.',
      logo: FreelancerLogo
    },
    {
      title: 'Senior Software Engineer',
      company: 'Cryptio',
      location: 'Paris',
      date: 'April - July 2023',
      type: 'Full-time',
      duration: '4 mos',
      description:
        'I led a team to transform digital asset data into auditable records for accounting, treasury, and tax filings, while enhancing my knowledge of relational databases and crypto asset accounting.',
      logo: CryptioLogo
    },
    {
      title: 'Fullstack Software Engineer',
      company: 'PayFit',
      location: 'Paris',
      date: 'Apr 2018 – Mar 2023',
      type: 'Full-time',
      duration: '5 yrs',
      description:
        'Developed internal tools automating government tasks and client data management. Built dashboards, implemented fuzzy search with virtualization for large datasets, and created event-driven CDC for real-time MongoDB-ElasticSearch sync. Transitioned to DevOps to streamline deployments, manage Kubernetes clusters, and implement infrastructure as code. Focused on improving scalability and monitoring services across the platform. Contributed to reducing system downtime and enhancing overall platform reliability. Mentored junior developers and participated in technical architecture decisions.',
      logo: PayfitLogo
    },
    {
      title: 'Fullstack Developer',
      company: 'Heek',
      location: 'Paris',
      date: 'Apr – Oct 2017',
      type: 'Full-time',
      duration: '6 mos',
      description:
        'I worked on a chatbot capable of generating websites according to customer needs. This role honed my skills in identifying and fixing critical bugs using debugging tools and provided valuable experience in team collaboration using agile methods. Additionally, it deepened my understanding of server-side rendering compared to client-side rendering.',
      logo: HeekLogo
    },
    {
      title: 'Mobile Application Developer',
      company: 'Export Japan',
      location: 'Osaka',
      date: 'May – Sep 2016',
      type: 'Internship',
      duration: '4 mos',
      description:
        'Implemented a BaaS server allowing multi-account management compatible with multiple providers and built a hybrid mobile app capable of utilizing smartphone hardware capabilities such as GPS, NFC, and vibrations.',
      logo: ExportJapanLogo
    }
  ];

  const education = [
    {
      title: 'Master in Computer Science',
      domain: 'Smart Systems',
      school: 'CY Tech',
      period: '2015 – 2017',
      location: 'Paris',
      description: `As the valedictorian of my engineering school, I had the privilege of returning to teach penetration testing and network security after graduating. During my time at the school, I worked on embedded systems and authored a well-received scientific article on an intelligent waste collection system, which explored the integration of smart systems for smart city applications. My curriculum included a robust set of teaching units related to computer science, providing a strong foundation for my career.`,
      image: 'cytech.jpg'
    },
    {
      title: 'B.E. in Computer Science',
      domain: 'Artificial Intelligence',
      school: 'Univ. Paris Cité',
      period: '2013 – 2015',
      location: 'Paris',
      description: `During my bachelor's studies, I focused on artificial intelligence and gained proficiency in programming languages such as Java and Python. My coursework included a comprehensive range of subjects, including algorithms, mathematics, and machine learning.`,
      image: 'paris_universite.jpg'
    }
  ];
</script>

<div class="flex items-center p-1cm">
  <div class="grid gap-4 grid-cols-3">
    <div class="space-y-4">
      <!-- Start Left Side -->
      <div class="shadow rounded-xl overflow-hidden">
        <!-- Start User Block -->
        <div class="h-24 bg-cover" style="background-image: url('cover.jpg')"></div>
        <div class="pt-14 p-7 bg-white relative">
          <img src="avatar.jpg" alt="Avatar" class="user-photo" />
          <div class="text-lg font-semibold mb-1.5">Thibault Vieux</div>
          <div class="text-sm text-gray-400 mb-1.5">Software Engineer</div>
          <ul class="flex space-x-4">
            {#each socials as { href, icon }}
              <li>
                <a {href} target="_blank" class="social-link-hover">
                  <i class={`bx ${icon} text-xl`}></i>
                </a>
              </li>
            {/each}
          </ul>
        </div>
      </div>
      <!-- End User Block -->
      <div class="p-5 block-section">
        <!-- Start Info Block -->
        <h2 class="block-title">Information</h2>
        <div class="space-y-4">
          <div class="flex justify-between">
            <div class="text-gray-400">📍 Location</div>
            <div class="font-medium text-right text-gray-600">Japan</div>
          </div>
          <div class="flex justify-between">
            <div class="text-gray-400">🛠️ Experience</div>
            <div class="font-medium text-right text-gray-600">6 years</div>
          </div>
          <div class="flex justify-between">
            <div class="text-gray-400">🍰 Birthday</div>
            <div class="font-medium text-right text-gray-600">1992/07/25</div>
          </div>
          <div class="flex justify-between">
            <div class="text-gray-400">📅 Available in</div>
            <div class="font-medium text-right text-gray-600">1 week</div>
          </div>
          <div class="flex justify-between">
            <div class="text-gray-400">🏠 Relocation</div>
            <div class="font-medium text-right text-gray-600">Amenable</div>
          </div>
          <div class="flex flex-col justify-between">
            <div class="flex content-center">
              <span class="bg-gray-100 text-gray-800 font-medium text-xs mr-1 px-1 py-1 rounded">
                🇫🇷 - FR Native
              </span>
              <span class="bg-gray-100 text-gray-800 font-medium text-xs mr-1 px-1 py-1 rounded">
                🇬🇧 - EN C1
              </span>
              <span class="bg-gray-100 text-gray-800 font-medium text-xs px-1 py-1 rounded">
                🇯🇵 - JP N4
              </span>
            </div>
          </div>
        </div>
      </div>
      <!-- End Info Block -->
      <div class="p-5 block-section flow-root">
        <!-- Start Skills Block -->
        <h2 class="block-title">Skills</h2>
        <div class="-m-2 flex flex-wrap">
          {#each skills as skill}
            <span class="skill-tag bg-purple-100 text-purple-800">{skill}</span>
          {/each}
        </div>
      </div>
      <!-- End Skills Block -->
      <div class="p-5 block-section flow-root">
        <!-- Start Projects Block -->
        <h2 class="block-title">Key Projects</h2>
        <div class="-m-2 flex flex-wrap">
          {#each projects as { href, name }}
            <a {href} target="_blank" class="skill-tag bg-blue-100 text-blue-800">{name}</a>
          {/each}
        </div>
      </div>
      <!-- End Projects Block -->
    </div>
    <!-- End Left Side -->

    <div class="space-y-3 col-span-2">
      <!-- Start Right Side -->
      <div class="p-5 block-section">
        <!-- Start About Me Block -->
        <h2 class="block-title">About me</h2>
        <p class="text-gray-600 text-sm">
          I specialize in <strong>building web applications</strong> from infrastructure deployment
          to frontend design and backend optimization. In my spare time, I program in Rust and work
          on home automation projects using my favorite web technologies. I also taught
          <b>penetration testing</b>
          and
          <b>network security</b> at university, gaining valuable experience and feedback. Living in
          Osaka since 2023, I stay active with hip-hop classes, daily gym workouts, and weekly
          drumming sessions. I am also learning Japanese by participating in language exchange events
          and previously attended a language school for 4 months.
        </p>
      </div>
      <!-- End About Me Block -->
      <div class="p-5 block-section">
        <!-- Start Experience Block -->
        <h2 class="block-title">Experience</h2>
        {#each experience as { company, date, description, duration, location, logo, title, type }, i}
          {@const SvelteComponent = logo}
          <div class={i === experience.length - 1 ? 'item-section' : 'mb-4 item-section'}>
            <SvelteComponent />
            <div class="w-full space-y-4">
              <div class="item-header">
                <div class="space-y-1.5">
                  <div class="font-medium">{title}</div>
                  <div class="flex space-x-4">
                    <div class="item-header-info">
                      <JobIcon />
                      <span>{company}</span>
                    </div>
                    <div class="item-header-info">
                      <MapIcon />
                      <span>{location}</span>
                    </div>
                  </div>
                </div>
                <div class="space-y-2 text-right">
                  <div class="job-item-badge">{type} · {duration}</div>
                  <div class="item-header-info">
                    <CalendarIcon />
                    <span>{date}</span>
                  </div>
                </div>
              </div>
              <p class="text-gray-600 text-sm" style="margin-top: 0.5rem">
                {description}
              </p>
              {#if i !== experience.length - 1}
                <div class="border-b border-gray-200"></div>
              {/if}
            </div>
          </div>
        {/each}
      </div>
      <!-- End Experience Block -->
      <div class="p-5 block-section">
        <!-- Start Education Block -->
        <h2 class="block-title">Education</h2>
        {#each education as { description, domain, image, location, period, school, title }, i}
          <div class={i === education.length - 1 ? 'item-section' : 'mb-4 item-section'}>
            <img alt="logo" class="rounded-md" src={image} height="48px" width="48px" />
            <div class="w-full space-y-4">
              <div class="item-header">
                <div class="space-y-1.5">
                  <div class="font-medium">{title}</div>
                  <div class="flex space-x-4">
                    <div class="item-header-info">
                      <EducationIcon />
                      <span>{school}</span>
                    </div>
                    <div class="item-header-info">
                      <MapIcon />
                      <span>{location}</span>
                    </div>
                  </div>
                </div>
                <div class="space-y-2 text-right">
                  <div class="job-item-badge">{domain}</div>
                  <div class="item-header-info">
                    <CalendarIcon />
                    <span>{period}</span>
                  </div>
                </div>
              </div>
              <p class="text-gray-600 text-sm" style="margin-top: 0.5rem">
                {description}
              </p>
              {#if i !== education.length - 1}
                <div class="border-b border-gray-200"></div>
              {/if}
            </div>
          </div>
        {/each}
      </div>
      <!-- End Education Block -->
    </div>
    <!-- End Right Side -->
  </div>
</div>
