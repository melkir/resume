<script>
  import CalendarIcon from '../components/icons/Calendar.svelte';
  import CryptioLogo from '../components/icons/Cryptio.svelte';
  import EducationIcon from '../components/icons/Education.svelte';
  import ExportJapanLogo from '../components/icons/ExportJapan.svelte';
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

  const skills = ['Rust', 'TypeScript', 'React', 'GraphQL', 'Docker', 'SQL', 'Git', 'MongoDB'];

  const projects = [
    { href: 'https://github.com/melkir/next-remote-gpio', name: 'Remote GPIO' },
    { href: 'https://github.com/melkir/hacking-workshop', name: 'Pentest Workshop' },
    { href: 'https://github.com/melkir/android-sms-push', name: 'SMS Push' },
    { href: 'https://github.com/melkir/API-Libraries-Demo', name: 'API Libraries Demo' }
  ];

  const experience = [
    {
      title: 'シニアエンジニア',
      company: 'Cryptio',
      location: 'パリ',
      date: '2023年4月 - 7月',
      type: 'フルタイム',
      duration: '4か月',
      description:
        'デジタル資産データを会計、財務、および税務申告のための監査可能な記録に変換する。',
      logo: CryptioLogo
    },
    {
      title: 'フルスタックエンジニア',
      company: 'PayFit',
      location: 'パリ',
      date: '2018年4月 - 2023年3月',
      type: 'フルタイム',
      duration: '5年間',
      description: '給与計算の自動化ソリューションの開発と保守。',
      logo: PayfitLogo
    },
    {
      title: 'フルスタック開発者',
      company: 'Heek',
      location: 'パリ',
      date: '2017年4月 - 10月',
      type: 'フルタイム',
      duration: '6か月',
      description: 'AIを使用したウェブサイトビルダーの開発。',
      logo: HeekLogo
    },
    {
      title: 'モバイル開発者',
      company: 'Export Japan',
      location: '大阪',
      date: '2016年5月 - 9月',
      type: 'インターンシップ',
      duration: '6か月',
      description:
        'スマートフォンのハードウェア機能を活用できるハイブリッドモバイルアプリを構築しました。',
      logo: ExportJapanLogo
    }
  ];

  const education = [
    {
      title: 'コンピュータサイエンス修士',
      domain: 'スマートシステム',
      school: 'CYテック',
      period: '2015 – 2017',
      location: 'パリ',
      image: 'cytech.jpg'
    },
    {
      title: 'コンピュータサイエンス学士',
      domain: '人工知能',
      school: 'パリ大学',
      period: '2013 – 2015',
      location: 'パリ',
      image: 'paris_universite.jpg'
    }
  ];
</script>

<div class="flex items-center sub-page">
  <div class="grid gap-4 grid-cols-3">
    <div class="space-y-4">
      <!-- Start Left Side -->
      <div class="shadow rounded-xl overflow-hidden">
        <!-- Start User Block -->
        <div class="h-24 bg-cover" style="background-image: url('cover.jpg')" />
        <div class="pt-14 p-7 bg-white relative">
          <img src="avatar.jpg" alt="Avatar" class="user-photo" />
          <div class="text-lg font-semibold mb-1.5">ティボ・ビュー</div>
          <div class="text-sm text-gray-400 mb-1.5">ソフトウェアエンジニア</div>
          <ul class="flex space-x-4">
            {#each socials as { href, icon }}
              <li>
                <a {href} target="_blank" class="social-link-hover">
                  <i class={`bx ${icon} text-xl`} />
                </a>
              </li>
            {/each}
          </ul>
        </div>
      </div>
      <!-- End User Block -->
      <div class="p-5 block-section">
        <!-- Start Info Block -->
        <h2 class="block-title">情報</h2>
        <div class="space-y-4">
          <div class="flex justify-between">
            <div class="text-gray-400">場所</div>
            <div class="font-medium text-right text-gray-600">日本</div>
          </div>
          <div class="flex justify-between">
            <div class="text-gray-400">経験</div>
            <div class="font-medium text-right text-gray-600">6年</div>
          </div>
          <div class="flex justify-between">
            <div class="text-gray-400">利用可能期間</div>
            <div class="font-medium text-right text-gray-600">1週間</div>
          </div>
          <div class="flex justify-between">
            <div class="text-gray-400">転居</div>
            <div class="font-medium text-right text-gray-600">転居可能</div>
          </div>
          <div class="flex justify-between">
            <div class="text-gray-400">言語</div>
            <div class="flex content-center">
              <span class="bg-gray-100 text-gray-800 font-medium text-xs mr-1 px-1 py-1 rounded">
                FR 🇫🇷
              </span>
              <span class="bg-gray-100 text-gray-800 font-medium text-xs mr-1 px-1 py-1 rounded">
                EN 🇬🇧
              </span>
              <span class="bg-gray-100 text-gray-800 font-medium text-xs px-1 py-1 rounded">
                JP 🇯🇵
              </span>
            </div>
          </div>
        </div>
      </div>
      <!-- End Info Block -->
      <div class="p-5 block-section flow-root">
        <!-- Start Skills Block -->
        <h2 class="block-title">スキル</h2>
        <div class="-m-2 flex flex-wrap">
          {#each skills as skill}
            <span class="skill-tag bg-purple-100 text-purple-800">{skill}</span>
          {/each}
        </div>
      </div>
      <!-- End Skills Block -->
      <div class="p-5 block-section flow-root">
        <!-- Start Projects Block -->
        <h2 class="block-title">主要プロジェクト</h2>
        <div class="-m-2 flex flex-wrap">
          {#each projects as { href, name }}
            <a {href} target="_blank" class="skill-tag bg-blue-100 text-blue-800">{name}</a>
          {/each}
        </div>
      </div>
      <!-- End Projects Block -->
    </div>
    <!-- End Left Side -->

    <div class="space-y-4 col-span-2">
      <!-- Start Right Side -->
      <div class="p-5 block-section">
        <!-- Start About Me Block -->
        <h2 class="block-title">自己紹介</h2>
        <p class="text-gray-600 text-sm">
          私はCryptioで<b>ウェブアプリケーションを構築</b
          >し、簡素化されたデジタル資産会計を専門としています。余暇にはRustでプログラムを書いたり、ホームオートメーションプロジェクトに取り組んでいます。また、大学で
          <b>ペネトレーションテスト</b>と<b>ネットワークセキュリティ</b
          >を教えており、貴重な経験とフィードバックを得ています。
        </p>
      </div>
      <!-- End About Me Block -->
      <div class="p-5 block-section">
        <!-- Start Experience Block -->
        <h2 class="block-title">経験</h2>
        {#each experience as { company, date, description, duration, location, logo, title, type }, i}
          <div class={i === experience.length - 1 ? 'item-section' : 'mb-4 item-section'}>
            <svelte:component this={logo} />
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
                <div class="border-b border-gray-200" />
              {/if}
            </div>
          </div>
        {/each}
      </div>
      <!-- End Experience Block -->
      <div class="p-5 block-section">
        <!-- Start Education Block -->
        <h2 class="block-title">アカデミー</h2>
        {#each education as { domain, image, location, period, school, title }, i}
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
              {#if i !== education.length - 1}
                <div class="border-b border-gray-200" />
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
